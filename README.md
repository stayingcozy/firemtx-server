<p align="center">
    <img src="assets/firemtx logo.png" height=300 width=300>
</p>

# firemtx-server
MediaMTX server for video streams with firebase communication. This media server acts as a relay from video source to viewer. 


### Why use a media server
Having a media server allows you to stream multiple video from your single source. For example one to the viewer and one to a Machine Learning / Computer Vision server. See pyRTCvision repository for ML server to process live video and upload results to your firebase database.

If desired to stream video to website (https) check out caddy-video-stream repository.

## Install
Install
* Mediamtx server
* golang
* firebase golang libraries

## Usage
Compile code. Run docker file to upload image to your favorite cloud provider's VM.
