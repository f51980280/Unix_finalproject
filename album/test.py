import numpy as np
import cv2
import os
import sys
import dlib
from skimage import io
from os import listdir
from os.path import isfile, isdir, join


detector = dlib.get_frontal_face_detector()
win = dlib.image_window()


total=0
face_count=0
mypath = 'picture/'
files = listdir(mypath)
face_cascade = cv2.CascadeClassifier('opencv/data/haarcascades/haarcascade_frontalface_default.xml')

for f in files:
  total+=1
  fullpath = join(mypath, f)
  origin_img = cv2.imread(fullpath)
  img = cv2.imread(fullpath)
  gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
  faces = face_cascade.detectMultiScale(gray)

  dets = detector(img)

  for (x,y,w,h) in faces:
    img2 = cv2.rectangle(img,(x,y),(x+w,y+h),(255,0,0),2)
    roi_gray = gray[y:y+h, x:x+w]
#    roi_color = img[y:y+h, x:x+w]


  for i, d in enumerate(dets):
        print("dets{}".format(d))
        print("Detection {}: Left: {} Top: {} Right: {} Bottom: {}"
            .format( i, d.left(), d.top(), d.right(), d.bottom()))
#  dets, scores, idx = detector.run(img, 1)
#  for i, d in enumerate(dets):
#        print("Detection {}, dets{},score: {}, face_type:{}".format( i, d, scores[i], idx[i]))
#  if len(faces)>=len(dets):
  print fullpath + ' has ' + str(len(faces)) + ' face(s) '
 # else:
#    print fullpath + ' has ' + str(len(dets)) + ' face(s) '
  #sp = img.shape #sp[0]=img.heigh    sp[1]=width  sp[2]=color
  img = cv2.resize(img, (256,256), interpolation=cv2.INTER_CUBIC)
  if len(faces)!=0:
    face_count+=1
#    win.clear_overlay()
#    win.set_image(img)
#    win.add_overlay(dets)
 #   dlib.hit_enter_to_continue()
    cv2.imwrite('detected/'+f, origin_img)
#  if len(faces)!=0:
    cv2.imshow('img',img)
    cv2.waitKey(0)
#    dlib.hit_enter_to_continue()
#    cv2.destroyAllWindows()
#  os.remove(fullpath)
print  'This web has ' + str(face_count*100/total)+'%  of picture includes face(s)'

cv2.destroyAllWindows()
