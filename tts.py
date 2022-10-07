import sys
import whisper

# import pyaudio

# FRAMES_PER_BUFFER = 3200
# FORMAT = pyaudio.paInt16
# CHANNELS = 1
# RATE = 16000
# p = pyaudio.PyAudio()
 
# # starts recording
# stream = p.open(
#    format=FORMAT,
#    channels=CHANNELS,
#    rate=RATE,
#    input=True,
#    frames_per_buffer=FRAMES_PER_BUFFER
# )
stream = sys.stdin.buffer
for chunk in stream:
    print(f'{chunk}')
    # text = float(chunk)
    # print(f'Processing Message from sys.stdin *****{text}*****')
