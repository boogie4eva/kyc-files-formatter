# kyc-files-formatter
Know your customer file formatter to remove unwanted characters from tonnes of files and storing them in another directory . <br>
This was initially written to solve a major problem with processing a large set of files with bad characters hosted in a directory . </br>
I developed this while i was working with Ntel to enable them to process a large set of files that had a funny character in it . </br>
This processes the large amount of files over 500Gb from a source directory and after processing each file writes the file to an output directory. </br>

```
 kyc-files-formatter -sourceDir=/files -destDir=./output 

```
