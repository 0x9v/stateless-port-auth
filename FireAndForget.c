#include<stdio.h>
#include<sys/socket.h>
#include<stdlib.h>
#include<arpa/inet.h>
#include<errno.h>
#include<string.h>

int main(int argc, char *argv[])
{
    char *end_ptr;
    char *ip;
    long port;

    if(argc != 3)
    {
        printf("Usage: %s <IP> <Port>\n", argv[0]);
        return 1;
    }//making sure that the argument count is correct

    ip = argv[1];
    
    //Validating the port input
    errno = 0;
    port = strtol(argv[2], &end_ptr, 10);

    if(errno == ERANGE)
    {
        printf("overflow hehe\n\n");
        return 1;
    }

    if(end_ptr == argv[2]||*end_ptr != '\0')
    {
        printf("invalid Char\n\n");
        return 1;
    }
    if(port<1 || port >65535)
    {
        printf("port out of range \n\n");
        return 1;
    }



    //Converting the ip addr from a string into a machine readable binary form
    struct in_addr addr;
    int res = inet_pton(AF_INET, ip, &addr);

    //validating the ip addr input
    if(res != 1)
    {
        printf("Nuh uh\n");
        return 1;
    }

    struct sockaddr_in target;
    memset(&target, 0, sizeof(target));//clean up our target

    target.sin_family = AF_INET;
    target.sin_port = htons(port);
    target.sin_addr = addr.s_addr;

    //incomplete code !!!
    return 0;
}
