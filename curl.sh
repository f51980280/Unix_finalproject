#!/bin/bash

printf "\E[0;36;40m"
echo ' 請輸入  (0)網址a  (1)imgur圖片網站 (2)其他圖片網址'
printf "\E[0m"
read input

#curl 'https://imgur.com/search?q=face' -o '#1.html'#imgur
#http://www.ttpaihang.com/vote/rank.php?voteid=1089&page=" #女優圖片
if [ $input = 1 ]
then
   curl 'https://imgur.com/search?q=face' -o '#1.html'
   grep -oh '//.*jpg' *.html >url2s.txt
   sed -e 's/\/\//https:\/\//p' url2s.txt > urls.txt
   sort -u urls.txt | wget -P picture -i-
elif [ $input = 0 ]
   then
       wget -nd -r -P picture -A jpg http://www.ttpaihang.com/vote/rank.php?voteid=1089
   else
       printf "\E[0;33;40m"
       echo "請輸入您要的網址"
       printf "\E[0m"
       read URL
       curl URL -o '#1.html'
       grep -oh '//.*jpg' *.html >url2s.txt
       sed -e 's/\/\//https:\/\//p' url2s.txt > urls.txt
       sort -u urls.txt | wget -P picture -i-
       wget -nd -r -P picture -A jpg URL
fi
rm picture/*.jpg.1
python test.py
