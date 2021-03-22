import os
import sys
import cv2
import json
import random
import multiprocessing
from imageai.Detection.Custom import CustomObjectDetection

execution_path = os.getcwd()


def simple_valid(model_path, json_path):
    parmas = sys.argv
    if len(parmas) <= 1:
        print('请输入要检测的类型，获取边界位置（box）或者判断目标是否存在（exist）')
        return

    if parmas[1] == 'exist' and len(parmas) <= 2:
        print('请输入目标label')
        return

    detector = CustomObjectDetection()
    detector.setModelTypeAsYOLOv3()
    detector.setModelPath(os.path.join(execution_path, model_path))
    detector.setJsonPath(os.path.join(execution_path, json_path))
    detector.loadModel()

    detections = detector.detectObjectsFromImage(
        input_image=os.path.join(execution_path, parmas[2]),
        output_image_path=os.path.join(execution_path, parmas[3])
    )

    if parmas[1] == 'box':
        info = [{'name': eachObject["name"], 'percent': str(eachObject["percentage_probability"]), 'box': eachObject["box_points"]} for eachObject in detections]
        json_info = json.dumps(info)
        print(json_info)
    elif parmas[1] == 'exist':
        exist = False
        for eachObject in detections:
            if eachObject["name"] == parmas[2]:
                exist = True
                break
        print(exist)

    # img = cv2.imread(os.path.join(execution_path, detected_pic_path))
    # cv2.namedWindow('image', 0)
    # cv2.resizeWindow('image', 1920, 1080)
    # cv2.imshow('image', img)
    # cv2.waitKey(0)
    # cv2.destroyAllWindows()


if __name__ == "__main__":
    model_path = "python/detection/detection_model-ex-072--loss-0000.256.h5"  # 训练模型路径
    json_path = "python/detection/detection_config.json"  # json文件路径

    p = multiprocessing.Process(target=simple_valid, args=(model_path, json_path))
    p.start()
    p.join()
