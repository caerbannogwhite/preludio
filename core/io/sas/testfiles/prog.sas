
/* NUMERIC */
LIBNAME testnum XPORT 'C:\Users\massi\Desktop\testnum.xpt';

DATA testnum.values;
  INPUT numvar1 numvar2 numvar3 numvar4;
  LABEL numvar1="A label for variable 1";
  DATALINES;
  1 1.1 -1.1 0.0
  16e62 16e-64 -999 999
  1.0000000000000001e+09 1.000000000000001e+09 1.00000000000001e+09 1.0000000000001e+09
  1.1e-05 1.2e-05 1.3e-05 1.4e-05
  1.1e-08 1.2e-08 1.3e-08 1.4e-08
  ;
RUN;

/* Close the library */
LIBNAME testnum CLEAR;


/* BINARY */
/* LIBNAME testbin XPORT 'C:\Users\massi\Desktop\testbin.xpt'; */
/*  */
/* DATA testbin.values; */
/*   LENGTH binaryvar1 8 binaryvar2 16 binaryvar3 32; */
/*   INPUT binaryvar1 binaryvar2 binaryvar3; */
/*   DATALINES; */
/*   255 65535 4294967295 */
/*   ; */
/* RUN; */
/*  */
/* Close the library */
/* LIBNAME testbin CLEAR; */


/* Date and time data types */
/*   INPUT datevar datetimevar;*/
/*   DATALINES;*/
/*   '16MAR2023'dt '16MAR2023:12:34:56'dt*/
/*   ;*/


/* CHARACTER */
LIBNAME testchar XPORT 'C:\Users\massi\Desktop\testchar.xpt';

DATA testchar.values;
  LENGTH charvar1 $10 charvar2 $20;
  INPUT charvar1 $ charvar2 $;
  DATALINES;
  abcdefghij abcdefghijklmnopqrst
  wbiwbui749 nionione983203jnfui2
  !"£$%^&*)( -_=+{}[]'@#~/?.>,<\|
  ;
RUN;

/* Close the library */
LIBNAME testchar CLEAR;
