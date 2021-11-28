import logging

import grpc
import random
 
import message_pb2, message_pb2_grpc

def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = message_pb2_grpc.CreateUserStub(channel)
        response = stub.Create(message_pb2.UserRequest(
            ID = "RandomID",
            Username="kalycodes",
            Firstname="Muhammad",
            Lastname="Abdullahi",
            DOB="January",
            Age=25,
            isAdmin=False
        ))
    print("UserCreated-------------- \n" + str(response))


if __name__ == '__main__':
    logging.basicConfig()
    run()

