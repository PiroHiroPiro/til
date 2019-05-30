/*
 *  fileinfo.c
 *  stat()を使ってファイルのプロパティを取得し、表示する
 */

#include <stdio.h>
#include <sys/types.h>
#include <sys/stat.h>

int main(int ac, char *av[]){
  struct stat info; //ファイル情報のためのバッファ
  if(ac > 1){
    if(stat(av[1], &info) != -1){
      show_stat_info(av[1], &info);
      return 0;
    }else{
      perror(av[1]);  //stat()エラーを報告する
    }
  }
  return 1;
}

show_stat_info(char *fname, struct stat *buf){
  /*
   *  名前＝値形式でstatの情報の一部を表示する
   */
   printf("   mode: %o\n", buf->st_mode); //タイプ＋モード
   printf("  links: %d\n", buf->st_nlink); //リンク数
   printf("   user: %d\n", buf->st_uid); //ユーザーID
   printf("  group: %d\n", buf->st_gid); //グループID
   printf("   size: %d\n", buf->st_size); //ファイルサイズ
   printf("modtime: %d\n", buf->st_mtime); //変更日時
   printf("   name: %s\n", fname); //ファイル名
}
