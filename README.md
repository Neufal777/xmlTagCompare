# Tag Check Duplicates [XML]

## Usage

```golang


files := []string{
	"file1.xml",
	"file2.xml",
	//you can add as many as you want
}

	//process files & get the map[string]int with all the duplicates
	res := FilesProcess(files, "ITEM_TAG_NAME")
	dups := CheckDuplicates(res)

	//print all duplicates
	ShowDuplicates(dups)

``` 

## Output

```sh
Processing..

url processed 20971

Comparing Urls...

https://www.URL.COM.com/15987565/508 Duplicated 15 Times
https://www.URL.OM.com/8415987565/43 Duplicated 12 Times
https://www.URL.COM.com/8415987565/98 Duplicated 8 Times

Total duplicates 35
2020/03/07 13:00:26 Exec Time 1.269446349s
```
