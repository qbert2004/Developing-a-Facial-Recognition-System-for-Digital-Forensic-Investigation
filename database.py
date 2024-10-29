import cv2
import face_recognition

classifiers = [

]

file_name = "hz.webp"

img = cv2.imread(file_name)

if img is None:
    print("Error: Image not loaded correctly.")
    exit()
 
height, width = img.shape[:2]

new_width = 700
new_height = int((new_width / width) * height)

img_resized = cv2.resize(img, (new_width, new_height))

gray_img_resized = cv2.cvtColor(img_resized, cv2.COLOR_BGR2GRAY)

rgb_img_resized = cv2.cvtColor(img_resized, cv2.COLOR_BGR2RGB)

face_locations = face_recognition.face_locations(rgb_img_resized)
for classifier in classifiers:
    face_cascade = cv2.CascadeClassifier(classifier)
    faces_resized = face_cascade.detectMultiScale(
        gray_img_resized, scaleFactor=1.1, minNeighbors=5, minSize=(40, 40)
    )

    if len(faces_resized) > 0:
        print(f"Detected {len(faces_resized)} face(s) with {classifier}.")


        for (x, y, w, h) in faces_resized:
            cv2.rectangle(img_resized, (x, y), (x + w, y + h), (0, 255, 0), 2)

if len(face_locations) > 0:
    print(f"Detected {len(face_locations)} face(s) with face_recognition.")

    for (top, right, bottom, left) in face_locations:
        cv2.rectangle(img_resized, (left, top), (right, bottom), (0, 255, 0), 2)



cv2.imshow('Detected Faces', img_resized)
cv2.waitKey(0)
cv2.destroyAllWindows()

output_path = "detected_faces_image.png"
cv2.imwrite(output_path, img_resized)