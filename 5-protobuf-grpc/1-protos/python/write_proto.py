from pb.my.my_pb2 import *

report = SellerParams()

first_seller = report.result.add()
first_seller.seller_id = 1
first_seller.rating = 11
first_seller.params["some_dobule"].double = 4.2

second_seller = report.result.add()
second_seller.seller_id = 3
second_seller.params["name"].string = "David"
second_seller.params["int"].int = 1


print("Report: ")
print(report)

with open('./output.bin', 'wb') as f:
    f.write(report.SerializeToString())
