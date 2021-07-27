# -*- coding:utf-8 -*-
import os
import yaml
import sys


current_path = os.path.abspath(os.path.dirname(__file__))
# print(current_path)

with open(current_path + '/config.yaml', 'r') as f:
    config = yaml.safe_load(f.read())
    overrides = config["network_overrides"]["constants"][config["selected_network"]]
    # 这里的overrides得到是config字典中匹配network_overrides节点下的constants的内容与selected_network下的内容。
    print("overrides:",overrides)
