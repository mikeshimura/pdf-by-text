# pdf-by-text
create complex pdf from text file.

[日本語の説明はこちら](https://github.com/mikeshimura/pdf-by-text/wiki/%E6%97%A5%E6%9C%AC%E8%AA%9E%E8%AA%AC%E6%98%8E)

This program is executable file for windows, mac or linux.  
Any language which create text file and execute child process can use this program.  

This program use github.com/signintech/gopdf for pdf generation.

##First part of following Simple report is as follows.  
Columns must be separated by tab(\t).

```go
P	mm	A4	L
F	IPAexG		14
C1	50.00	15.00	Sales Report
F	IPAexG		12
C1	240.00	20.00	page
C1	260.00	20.00	1
C1	15.00	23.00	D No
C1	30.00	23.00	Dept
C1	60.00	23.00	Order
C1	90.00	23.00	Stock
C1	120.00	23.00	Name
CR	135.00	23.00	25.00	Unit Price
CR	160.00	23.00	20.00	Qty
CR	190.00	23.00	20.00	Amount
```



Simple report sample
![Simple1](https://bytebucket.org/mikeshimura/goreport/wiki/image/simple1.jpg "Simple1")  
[text](https://bytebucket.org/mikeshimura/goreport/wiki/pdf-by-text/simple1.txt)
[pdf](https://bytebucket.org/mikeshimura/goreport/wiki/pdf-by-text/simple1.pdf)  


Medium report sample
![Medium1](https://bytebucket.org/mikeshimura/goreport/wiki/image/medium1.jpg "Medium1")

[text](https://bytebucket.org/mikeshimura/goreport/wiki/pdf-by-text/medium1.txt)
[pdf](https://bytebucket.org/mikeshimura/goreport/wiki/pdf-by-text/medium1.pdf)  



Complex report sample
![Complex1](https://bytebucket.org/mikeshimura/goreport/wiki/image/complex1.jpg "Complex1")

[text](https://bytebucket.org/mikeshimura/goreport/wiki/pdf-by-text/complex1.txt)
[pdf](https://bytebucket.org/mikeshimura/goreport/wiki/pdf-by-text/complex1.pdf)  


![Complex2](https://bytebucket.org/mikeshimura/goreport/wiki/image/complex2.jpg "Complex2")

[text](https://bytebucket.org/mikeshimura/goreport/wiki/pdf-by-text/complex2.txt)
[pdf](https://bytebucket.org/mikeshimura/goreport/wiki/pdf-by-text/complex2.pdf)  

##Execution  

pdf-by-text(.exe) -f fontfile -o outputfile inputfile

fontfile default is font.yml  
outputfile default is inputfile name with file extension changed to pdf

##Download  
[windows 386](https://bytebucket.org/mikeshimura/goreport/wiki/pdf-by-text/pdf-by-text_windows_386.exe)  
[windows amd64](https://bytebucket.org/mikeshimura/goreport/wiki/pdf-by-text/pdf-by-text_windows_amd64.exe)  
[mac 386](https://bytebucket.org/mikeshimura/goreport/wiki/pdf-by-text/pdf-by-text_darwin_386)  
[mac amd64](https://bytebucket.org/mikeshimura/goreport/wiki/pdf-by-text/pdf-by-text_darwin_386)  
[linux 386](https://bytebucket.org/mikeshimura/goreport/wiki/pdf-by-text/pdf-by-text_linux_386)  
[linux amd64](https://bytebucket.org/mikeshimura/goreport/wiki/pdf-by-text/pdf-by-text_linux_amd64)  

##concept
- Simple command to draw object
- Support unit of mm(millimeter), pt(point) and in(Inch).
- Line can be used as gray fill pattern.  

##font file  
```
-
  name:IPAexG
  file:ttf\ipaexg.ttf
-
  name:MPBOLD
  file:ttf\mplus-1p-bold.ttf
```
##Draw commands  

Elements must be separated by tab(\t).

- Page Setup  
 - Predefined page  
   A {unit} {size} {orientation}  
   //unit mm(millimeter), pt(point), in(Inch)  
   //size A4 or LTR
   //orientation P(portrate) or L(Landscape)  
 - Frexible Page  
   A1 {unit} {width} {height}  


- Font setting  
 - Font face
  F {fontname} {size} {style}  
  //style U for underline
 - Text Color  
  TC {red}  {green}  {blue}  
  //color 255 - 0  
 - Text Gray Scale  
  GF {scale}  
  // scale 100 - 0  


- Text Draw  
 - Normal text  
  C1 {x} {y} {content}  
 - Right Jusify text  
  CR {x} {y} {width} {content}  


- Line Draw  
 - Line Type  
   LT {linetype} {width}  
   //lineType "dashed" ,"dotted","straight" ""="straight"   
 - Gray Stroke (Gray Scale)  
  GS  {scale}  
  // scale 100 - 0  
 - Horizontal Line  
   LH {x1} {y} {x2}    
 - Vertical Line  
   LV {x} {y1} {y2}
 - Line    
   L {x1} {y1} {x2} {y2}


- Shape Draw  
  - Rectangle  
   R {x1} {y1} {x2} {y2} {line width}  
  - Oval  
   O {x1} {y1} {x2} {y2}  


- Image Draw  
 I {path} {x1} {y1} {x2} {y2}

##Limitation  
- Font style not allow B(bold) and I(italic).
- Line, Rect and Oval are Black and Gray only.
- Image file format is jpeg only.  
