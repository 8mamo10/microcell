import time
import RPi.GPIO as GPIO

GPIO.setmode(GPIO.BCM)
GPIO.setup(4, GPIO.OUT)

p = GPIO.PWM(4, 50)

p.start(0.0)

dc = 2.5
while True:
    p.ChangeDutyCycle(dc)
    time.sleep(1)
    if dc == 2.5:
        dc = 12
    else:
        dc = 2.5
