#!/bin/bash

PAR_NUM=$#

#命名規則為DIRNAME-DATE-TIME.tgz
DIRNAME=$(dirname `pwd`)
DATE=`date +%Y%m%d`
TIME=`date +%H%M`
BACKNAME=$DIRNAME-$DATE-$TIME.tgz
USER=`whoami`
PWD=`pwd`
#如果沒有此資料夾就創建
if [ ! -d /home/$USER/backup ]
then	
	mkdir /home/$USER/backup
fi
#將BACKNAME中     ##:左邊數來最後一個   */: / 的左邊的字串全部刪除
NAME=${BACKNAME##*/}

if [ $PAR_NUM = 0 ]
then
	echo "沒給參數，直接壓縮"
	#如果沒給參數就直接壓縮
	tar zcvf  /home/$USER/backup/$NAME *
elif [ $PAR_NUM = 1 ]
then
	echo  1個參數
	 #如果剛好一個參數那就判斷是-n還是-q還是都不是
	PAR=$1
	FIR_PAR=${PAR:0:1}
	INT_PAR=${PAR:1}
	if [ "$PAR" = "-q" ]
	then
		echo "刪除排程"
		crontab -r
	elif [ $FIR_PAR = "-" ]&&[ $INT_PAR -gt 0 ]
	then
		echo 將在每 $INT_PAR 分鐘備份
		tar zcvf  /home/$USER/backup/$NAME *
		CRONTAB_BACKUP="*/$INT_PAR * * * * $PWD/mybackup.sh"
		(crontab -u $USER -l; echo "$CRONTAB_BACKUP" )| crontab -u $USER -
		exit
		
	else
		echo "指令只能是./mybackup 或 ./mybackup -n (n是int) 或 ./mybackup -q"
	fi		
else
	echo 你給了 $PAR_NUM 個參數
	echo "最多給1個參數"
fi

