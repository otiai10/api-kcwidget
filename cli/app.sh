#!/bin/sh

print_help()
{
  echo "\tstart\t起動"
  echo "\tstop\t終了"
  echo "\tstate\t状態確認"
  exit 0
}

# script for run the daemon
if [ $# -lt 1 ]; then
  print_help
fi

cur_date=`date '+%Y.%m%d'`
log_path='log/'$cur_date'.log'

case $1 in
  "start" )
    nohup python $PWD/minimum.py >> $log_path &
    sleep 1s
    break;;
  "stop" )
    ps aux | grep python | grep $PWD/minimum.py | grep -v grep | awk '{print $2}' | xargs kill -9
    sleep 1s
    break;;
  "state" )
    ps aux | grep python | grep $PWD/minimum.py | grep -v grep
    break;;
  "help" )
    echo "\033[0;33m"
    print_help
    break;;
  * )
    print_help
esac
