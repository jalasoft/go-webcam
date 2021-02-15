#include "v4l2-binding.h"
#include<sys/ioctl.h>
#include<sys/mman.h>
#include<string.h>
#include<stdio.h>
#include<errno.h>

struct v4l2_capability *queryCapability(int fd) {
    struct v4l2_capability* cap = malloc(sizeof(struct v4l2_capability));
    ioctl(fd, VIDIOC_QUERYCAP, cap);
    return cap;
}

struct v4l2_fmtdesc* queryFormats(int fd, struct v4l2_fmtdesc* desc) {
    if (desc == NULL) {
        desc = malloc(sizeof(struct v4l2_fmtdesc));
        desc->index = 0;
        desc->type = V4L2_BUF_TYPE_VIDEO_CAPTURE;
    } else {
        desc->index++;
    }

    int r = ioctl(fd, VIDIOC_ENUM_FMT, desc);

    if (errno == EINVAL) {
        free(desc);
        return NULL;
    }

    return desc;    
}

struct v4l2_frmsizeenum* queryFramesizes(int fd, __u32 pixformat, struct v4l2_frmsizeenum* info) {

    if (info == NULL) {
        info = malloc(sizeof(struct v4l2_frmsizeenum));
        info->index = 0;
        info->pixel_format = pixformat;
    } else {
        info->index++;
    }

    int result = ioctl(fd, VIDIOC_ENUM_FRAMESIZES, info);

    if (result < 0) {
        free(info);
        return NULL;
    }

    return info;
}

void setDiscreteFrameSize(int fd, __u32 pixformat, __u32 width, __u32 height) {

    struct v4l2_format format;
    format.type = V4L2_BUF_TYPE_VIDEO_CAPTURE;
    format.fmt.pix.pixelformat = pixformat;
    format.fmt.pix.width = width;
    format.fmt.pix.height = height;

    ioctl(fd, VIDIOC_S_FMT, &format);
}

void requestBuffer(int fd) {

    struct v4l2_requestbuffers request;
    request.type = V4L2_BUF_TYPE_VIDEO_CAPTURE;
    request.memory = V4L2_MEMORY_MMAP;
    request.count = 1;

    ioctl(fd, VIDIOC_REQBUFS, &request);

}

struct v4l2_buffer* queryBuffer(int fd) {
    struct v4l2_buffer* buffer = malloc(sizeof(struct v4l2_buffer));
    memset(buffer, 0, sizeof(struct v4l2_buffer));

    buffer->type = V4L2_BUF_TYPE_VIDEO_CAPTURE;
    buffer->memory = V4L2_MEMORY_MMAP;
    buffer->index = 0;

    ioctl(fd, VIDIOC_QUERYBUF, buffer);
	
    return buffer;
}

void* mmap2(int fd, struct v4l2_buffer* buf) {

    void* buff_start = mmap(NULL, buf->length, PROT_READ | PROT_WRITE, MAP_SHARED, fd, buf->m.offset);

    if (buff_start == MAP_FAILED) {
        return 0;
    }

    memset(buff_start, 0, buf->length);
    return buff_start;
}

void munmap2(void* mem_adr, uint length) {
    munmap(mem_adr, length);   
}

struct v4l2_buffer* newBuffer() {

    struct v4l2_buffer* buffer = malloc(sizeof(struct v4l2_buffer));
    buffer->type = V4L2_BUF_TYPE_VIDEO_CAPTURE;
    buffer->memory = V4L2_MEMORY_MMAP;
    buffer->index = 0;

    return buffer;
}

void queueBuffer(int fd, struct v4l2_buffer* buff) {
    ioctl(fd, VIDIOC_QBUF, buff);
}

void dequeueBuffer(int fd, struct v4l2_buffer* buff) {
    ioctl(fd, VIDIOC_DQBUF, buff);
}

void streamOn(int fd) {
    __u32 type = V4L2_BUF_TYPE_VIDEO_CAPTURE;
    ioctl(fd, VIDIOC_STREAMON, &type);
}

void streamOff(int fd) {
    __u32 type = V4L2_BUF_TYPE_VIDEO_CAPTURE;
    ioctl(fd, VIDIOC_STREAMOFF, &type);
}