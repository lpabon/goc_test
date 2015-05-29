#ifndef _CHANNEL_H
#define _CHANNEL_H

typedef struct {
    int val;
} msg_t;

typedef struct {
    pthread_cond_t send_c;
    pthread_cond_t recv_c;
    pthread_mutex_t m;
    msg_t *msg;
} channel_t;

void ch_init(channel_t *ch); 
void ch_send(channel_t *ch, msg_t *msg);
msg_t * ch_recv(channel_t *ch);

#endif
