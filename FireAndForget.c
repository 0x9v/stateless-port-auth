#include <netinet/in.h>
#include<stdio.h>
#include<sys/socket.h>
#include<stdlib.h>
#include<arpa/inet.h>

int main(int argc, char *argv[])
{
    char *ip;
    int port;

    if(argc != 3)
    {
        printf("Usage: %s <IP> <Port>\n", argv[0]);
        return 1;
    }//making sure that the argument count is correct

    ip = argv[1];
    port = atoi(argv[2]);
    printf("IP= %s, port = %d /n",ip, port);//i was testing the arguments here, i'll remove ts bullocks later

    //Converting the ip addr from a string into a machine readable binary form
    struct in_addr adrr;
    inet_pton(AF_INET, ip, &adrr);

    //INCOMPLETE CODE!!!
    return 0;
}