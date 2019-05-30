/*
 *  who2.c
 *  utmpファイルをオープンし、読み出した結果を表示する。
 *  /etc/utmpを読み出し、その中の情報のリストを表示する。
 *  空レコードが出力されないようにし、時刻を適切に整形する。
 */

#include <stdio.h>
#include <utmp.h>
#include <fcntl.h>
#include <unistd.h>
#include <time.h>

#define SHOWHOST  //出力にリモートマシン名を組み込む
//ut_typeの定義
#define EMPTY         0
#define RUN_LVL       1
#define BOOT_TIME     2
#define OLD_TIME      3
#define NEW_TIME      4
#define INIT_PROCESS  5 //"init"によって起動されたプロセス
#define LOGIN_PROCESS 6 //ログイン待ちの"getty"プロセス
#define USER_PROCESS  7 //ユーザープロセス
#define DEAD_PROCESS  8

void show_info(struct utmp *);
void showtime(long);

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
 *  レコードにユーザー名が含まれていない場合には何も表示しない
 */

void show_info(struct utmp *utbufp){
  if(utbufp->ut_type != USER_PROCESS){  //ユーザーのみ
    return;
  }

  ptintf("%-8.8s", utbufp->ut_name); //ログイン名
  ptintf(" "); //スペース
  ptintf("%-8.8s", utbufp->ut_line); //tty
  ptintf(" "); //スペース
  showtime(utbufp->ut_time); //ログイン時刻
  #ifdef SHOWHOST
    if(utbufp->ut_host[0] != '\0'){
      ptintf("(%s)", utbufp->ut_host); //ホスト
    }
  #endif
    ptintf("\n"); //改行
}

void showtime(long timeval) {
 /*
  * 人間が理解しやすい形式で時刻を表示する。
  * ctimeを使って文字列を組み立ててから、その一部を抜き出す。
  * 注意: %12.12sは文字列をchar12個分で出力し、長さを12バイト以下に制限する。
  */

  char *cp;  //時刻のアドレスを保持するポインタ
  cp = ctime(&timeval); //時刻を文字列に変換する　Mon Feb 4 00:46:40 EST 1991 0123456789012345.
  printf("%12.12s", cp+4);  //位置4から12文字分を抜き出す
}
