#include<stdio.h>
#include<sys/socket.h>
#include<stdlib.h>
#include<arpa/inet.h>
#include<errno.h>
#include<string.h>
#include<unistd.h>

//Hermes - UDP Packet Sender

void print_usage(char *prog)
{
    printf("\n\nUsage: %s <IP> <PORT> <PAYLOAD>\n\n", prog);
    printf("sends a UDP packet to the specified IP and port\n");
    printf("\nArguments: \n");
    printf("IP: Target IPv4 adress (e.g. 127.0.0.1)\n");
    printf("PORT: Target port (1 - 65535)\n");
    printf("PAYLOAD: Message to send\n\n");

}

int main(int argc, char *argv[])
{
    char *end_ptr = NULL;
    char *ip;
    long port;
    char *payload;
    size_t string_size;

    if(argc != 4)
    {
        print_usage(argv[0]);
        return 1;
    }//making sure that the argument count is correct

    ip = argv[1];
    
    //Validating the port input
    errno = 0;
    port = strtol(argv[2], &end_ptr, 10);

    if(errno == ERANGE)
    {
        return 2;//Overflow Error code
    }

    if(end_ptr == argv[2]||*end_ptr != '\0')
    {
        return 3;//garbage input Error code
    }
    if(port<1 || port >65535)
    {
        return 4;//Invalid PORT Error code
    }

    payload = argv[3];
    string_size = strlen(payload);

    //Converting the ip addr from a string into a machine readable binary form and storing it in "addr"
    struct in_addr addr;
    int res = inet_pton(AF_INET, ip, &addr);

    //validating the ip address input
    if(res != 1)
    {
        return 5;
    }

    //Declaring A variable of type struct sockaddr to store the target's data
    struct sockaddr_in target;

    //Cleaning up "target"
    memset(&target, 0, sizeof(target));

    //Assigning the data
    target.sin_family = AF_INET;
    target.sin_port = htons(port);
    target.sin_addr.s_addr = addr.s_addr;

    //Creating a socket
    int sock = socket(AF_INET, SOCK_DGRAM, 0);
    
    //Verifying if the socket succeeded
    if(sock<0)
    {
        return 6;//Socket Failed code
    }

    //Sending our packet
    ssize_t sent = sendto(sock, payload, string_size, 0, (struct sockaddr*)&target, sizeof(target));
    
    //Verifying that the packet was sent
    if(sent<0)
    {
        return 7;
    }

    //Closing our socket
    close(sock);


    return 0;
}

//GumWub