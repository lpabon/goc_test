#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <pthread.h>
#include "channel.h"

void
ch_init(channel_t *ch) {
    ch->msg = NULL;
    pthread_mutex_init(&ch->m, NULL);
    pthread_cond_init(&ch->send_c, NULL);
    pthread_cond_init(&ch->recv_c, NULL);
}

void
ch_send(channel_t *ch, msg_t *msg) {

    pthread_mutex_lock(&ch->m);
    while (NULL != ch->msg) {
        pthread_cond_wait(&ch->send_c, &ch->m);
    }

    ch->msg = msg;

    pthread_cond_signal(&ch->recv_c);

    // Wait until recv has it
    while (NULL != ch->msg) {
        pthread_cond_wait(&ch->send_c, &ch->m);
    }
    pthread_mutex_unlock(&ch->m);
}

msg_t *
ch_recv(channel_t *ch) {

    msg_t *msg;

    pthread_mutex_lock(&ch->m);
    while (NULL == ch->msg) {
        pthread_cond_wait(&ch->recv_c, &ch->m);
    }

    msg = ch->msg;
    ch->msg = NULL;

    pthread_cond_signal(&ch->send_c);
    pthread_mutex_unlock(&ch->m);

    return msg;
}

