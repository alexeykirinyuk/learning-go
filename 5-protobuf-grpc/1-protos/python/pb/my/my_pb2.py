# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: my/my.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0bmy/my.proto\x12\x12golang.school.demo\"y\n\nParamValue\x12\x18\n\x06\x64ouble\x18\x01 \x01(\x01H\x00R\x06\x64ouble\x12\x12\n\x03int\x18\x02 \x01(\x03H\x00R\x03int\x12\x14\n\x04\x62ool\x18\x03 \x01(\x08H\x00R\x04\x62ool\x12\x18\n\x06string\x18\x04 \x01(\tH\x00R\x06stringB\r\n\x0bvalue_oneof\"\xb1\x02\n\x0cSellerParams\x12=\n\x06result\x18\x01 \x03(\x0b\x32%.golang.school.demo.SellerParams.ItemR\x06result\x1a\xe1\x01\n\x04Item\x12\x1b\n\tseller_id\x18\x01 \x01(\x03R\x08sellerId\x12\x16\n\x06rating\x18\x02 \x01(\x01R\x06rating\x12I\n\x06params\x18\x03 \x03(\x0b\x32\x31.golang.school.demo.SellerParams.Item.ParamsEntryR\x06params\x1aY\n\x0bParamsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x34\n\x05value\x18\x02 \x01(\x0b\x32\x1e.golang.school.demo.ParamValueR\x05value:\x02\x38\x01\x42\x30Z.github.com/alexeykirinyuk/learning-go/protos-1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'my.my_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z.github.com/alexeykirinyuk/learning-go/protos-1'
  _SELLERPARAMS_ITEM_PARAMSENTRY._options = None
  _SELLERPARAMS_ITEM_PARAMSENTRY._serialized_options = b'8\001'
  _PARAMVALUE._serialized_start=35
  _PARAMVALUE._serialized_end=156
  _SELLERPARAMS._serialized_start=159
  _SELLERPARAMS._serialized_end=464
  _SELLERPARAMS_ITEM._serialized_start=239
  _SELLERPARAMS_ITEM._serialized_end=464
  _SELLERPARAMS_ITEM_PARAMSENTRY._serialized_start=375
  _SELLERPARAMS_ITEM_PARAMSENTRY._serialized_end=464
# @@protoc_insertion_point(module_scope)