#include<stdlib.h>
#include<linux/videodev2.h>

struct frmsize_node {
    struct v4l2_frmsizeenum* value;
    struct frmsize_node* next;
};

struct v4l2_capability *queryCapability(int fd);

struct v4l2_fmtdesc* queryFormats(int fd, struct v4l2_fmtdesc* desc);

struct v4l2_frmsizeenum* queryFramesizes(int fd, __u32 fixformat, struct v4l2_frmsizeenum* info);

void setDiscreteFrameSize(int fd, __u32 pixformat, __u32 width, __u32 height);

void requestBuffer(int fd);

struct v4l2_buffer* queryBuffer(int fd);

void* mmap2(int fd, struct v4l2_buffer* buff);

void munmap2(void* mem_adr, uint length);

struct v4l2_buffer* newBuffer();

void queueBuffer(int fd, struct v4l2_buffer* buff);

void dequeueBuffer(int fd, struct v4l2_buffer* buff);

void streamOn(int fd);

void streamOff(int fd);