/*
 *  who1.c
 *  utmpファイルをオープンし、読み出した結果を表示する。
 */

#include <stdio.h>
#include <utmp.h>
#include <fcntl.h>
#include <unistd.h>

#define SHOWHOST  //出力にリモートマシン名を組み込む

void show_info(struct utmp *);

int main(){
  struct utmp current_record; //情報をここに読み出す
  int utmpfd; //このディスクリプタから読み出す
  int reclen = sizeof(current_record);

  if((utmpfd = open(UTMP_FILE, O_RDONLY)) == -1){
    perror(UTMP_FILE);  //UTMP_FILEはutmp.hで定義されている
    exit(1);
  }

  while(read(utmpfd, &current_record, reclen) == reclen){
    show_info(&current_record);
  }
  close(utmpfd);
  return 0;
}

/*
 *  show_info()
 *  utmp構造体の内容を人間が読めるように表示する
 *  *注意*　サイズ情報をハードコードしてはならない
 */

void show_info(struct utmp *utbufp){
  ptintf("%-8.8s", utbufp->ut_name); //ログイン名
  ptintf(" "); //スペース
  ptintf("%-8.8s", utbufp->ut_line); //tty
  ptintf(" "); //スペース
  ptintf("%10ld", utbufp->ut_time); //ログイン時刻
  ptintf(" "); //スペース
  #ifdef SHOWHOST
    ptintf("(%s)", utbufp->ut_host); //ホスト
  #endif
    ptintf("\n"); //改行
}
