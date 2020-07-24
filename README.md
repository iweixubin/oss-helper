怎么识别文件为同一个文件

 * 文件长度，拥有多少个 byte  
 * 取文件尾部 8个byte  
 * 文件的 MD5  
 * 文件的 SHA512  
 * 文件的扩展名  


如果两个文件长度不一样，那么一定不是同一个文件。

为什么取文件尾部 8个byte，作为一个值，是在长度相同的基础上再一次作为区分，  
为什么不取头部，因为文件头部有文件的标识，所以很大概率是相同的，  
为什么有 MD5 和 SHA512 还要取尾部，因为打算只用 长度+尾部 作为数据库索引。

为什么要 MD5 和 SHA512 两个，减少碰撞概率。

为什么要文件扩展名，  
因为文件标识 zip 可以对应的文件扩展名有： docx、apk、epub...
同一个文件的扩展名可以改来改去，但作为最终展示，只能按用户的来~



# 如何识别文件的真正格式

由于文件扩展名(后缀名)是可以随意修改的，所以通过扩展名来判断是不可靠的~

至于怎么识别如何识别文件的真正格式，  
我们可以读取写在文件头部的文件标识(signature)来识别，  
文件的特征，请浏览 [维基百科 List of file signatures](https://en.wikipedia.org/wiki/List_of_file_signatures)  
怎么读取文件标识，可以查看 GO 源码：https://github.com/golang/go/blob/master/src/net/http/sniff.go  


下文是简介：

## 精准识别

使用 字节模式（byte pattern）来识别~

### [PNG 的文件结构](https://zh.wikipedia.org/wiki/PNG)

PNG 图像格式文件由一个8字节的PNG文件标识（file signature）域和3个以上的后续数据块（IHDR、IDAT、IEND）组成。

PNG 文件包括 8字节文件署名（89 50 4E 47 0D 0A 1A 0A，十六进制），用来识别 PNG 格式。

<table>
    <tr>
        <th>十六进制</th>
        <th>含义</th>
    </tr>
    <tr>
        <td>89</td>
        <td>用于检测传输系统是否支持8位的字符编码（8 bit data），<br />
            用以减少将文本文件被错误的识别成PNG文件的机会，反之亦然。</td>
    </tr>
    <tr>
        <td>50 4E 47</td>
        <td>PNG每个字母对应的ASCII，让用户可以使用文本编辑器查看时，识别出是PNG文件。</td>
    </tr>
    <tr>
        <td>0D 0A</td>
        <td>DOS风格的换行符（CRLF）。用于DOS-Unix数据的换行符转换。</td>
    </tr>
    <tr>
        <td>1A</td>
        <td>在DOS命令行下，用于阻止文件显示的文件结束符。</td>
    </tr>
    <tr>
        <td>0A</td>
        <td>Unix风格的换行符（LF）。用于Unix-DOS换行符的转换。</td>
    </tr>
</table>

也就是说，如果我们读取一个文件的前 8个字节，  
内容为 [89 50 4E 47 0D 0A 1A 0A]，那么该文件就是 PNG 文件。

### [GIF 的文件结构](https://en.wikipedia.org/wiki/GIF)

```
byte#  hexadecimal  text or
(hex)               value       Meaning
0:     47 49 46
       38 39 61     GIF89a      Header
                                Logical Screen Descriptor
6:     03 00        3            - logical screen width in pixels
8:     05 00        5            - logical screen height in pixels
A:     F7                        - GCT follows for 256 colors with resolution 3 × 8 bits/primary;
                                   the lowest 3 bits represent the bit depth minus 1, the highest true bit means that the GCT is present
B:     00           0            - background color #0
C:     00                        - default pixel aspect ratio
                   R    G    B  Global Color Table
D:     00 00 00    0    0    0   - color #0 black
10:    80 00 00  128    0    0   - color #1
 :                                       :
85:    00 00 00    0    0    0   - color #40 black
 :                                       :
30A:   FF FF FF  255  255  255   - color #255 white
30D:   21 F9                    Graphic Control Extension (comment fields precede this in most files)
30F:   04           4            - 4 bytes of GCE data follow
310:   01                        - there is a transparent background color (bit field; the lowest bit signifies transparency)
311:   00 00                     - delay for animation in hundredths of a second: not used
313:   10          16            - color #16 is transparent
314:   00                        - end of GCE block
315:   2C                       Image Descriptor
316:   00 00 00 00 (0,0)         - NW corner position of image in logical screen
31A:   03 00 05 00 (3,5)         - image width and height in pixels
31E:   00                        - no local color table
31F:   08           8           Start of image - LZW minimum code size
320:   0B          11            - 11 bytes of LZW encoded image data follow
321:   00 51 FC 1B 28 70 A0 C1 83 01 01
32C:   00                        - end of image data
32D:   3B                       GIF file terminator
```


## 匹配识别

从 [维基百科 List of file signatures](https://en.wikipedia.org/wiki/List_of_file_signatures) 可以看到，  
avi 文件标识是这样的一种模式：`52 49 46 46 ?? ?? ?? ?? 41 56 49 20`

要识别这种文件，可以使用 [pattern mask](https://mimesniff.spec.whatwg.org/#matching-an-audio-or-video-type-pattern)


## 内容识别

像 html，xml，json 等，我们是通过起内容来识别的，  
毕竟它们都可以通过在 txt 基础上修改扩展名就可以了~ 