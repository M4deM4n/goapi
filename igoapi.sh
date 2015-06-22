#!/bin/bash

APP_NAME="GoAPI"
APP_FILENAME="goapi"
APP_PID="/var/run/$APP_FILENAME.pid"
APP_PATH="/home/ec2-user/go/src/github.com/m4dem4n/goapi"

TMP_FILE="/tmp/status_$APP_FILENAME"

start() {
	checkpid
	STATUS=$?
	echo -n "Starting $APP_NAME: "
	
	if [ $STATUS -ne 0 ]; then
		nohup $APP_PATH/$APP_FILENAME > $APP_PATH/$APP_FILENAME.out 2> $APP_PATH/$APP_FILENAME.err < /dev/null &
		echo $! > $APP_PID
		echo -e "[\\033[1;32mSuccess\\033[0;39m]"
	else
		echo -e "[\\033[1;31mFailure\\033[0;39m]"
	fi
}


stop() {
	checkpid
	STATUS=$?
	echo -n "Stopping $APP_NAME: "

	if [ $STATUS -eq 0 ]; then
		kill `cat $APP_PID`
		rm $APP_PID
		echo -e "[\\033[1;32mSuccess\\033[0;39m]"
	else
		echo -e "[\\033[1;31mFailure\\033[0;39m]"
	fi
}


checkpid() {
    STATUS=9
    
    if [ -f $APP_PID ] ;
	then
		#echo "Is Running if you can see next line with $APP_NAME"
		ps -Fp `cat $APP_PID` | grep $APP_FILENAME > $TMP_FILE
		if [ -f $TMP_FILE -a -s $TMP_FILE ] ;
			then
				STATUS=0
				#"Is Running (PID `cat $APP_PID`)"
			else
				STATUS=2
				#"Stopped incorrectly"
			fi
		
		## Clean after yourself	
		rm -f $TMP_FILE
	else
		STATUS=1
		#"Not Running"
	fi
	
	return $STATUS
}

case "$1" in
	'start')
		start
		;;
	'stop')
		stop
		;;
	'restart')
		stop
		start
		;;
	*)
		echo "Usage: $0 { start | stop | restart }"
		exit 1
		;;
esac

exit 0
