CC = gcc
CFLAGS = -c -Wall \
	-I/usr/include/php/20151012 \
	-I/usr/include/php/20151012/Zend \
	-I/usr/include/php/20151012/TSRM \
	-I/usr/include/php/20151012/sapi \
	-I/usr/include/php/20151012/main
LDFLAGS = -lphp7 

all: 
	$(CC) -o embed_test.o embed_test.c $(CFLAGS)
	$(CC) -o embed_test embed_test.o $(LDFLAGS)
  
