# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: savourrpc/airdrop.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from savourrpc import common_pb2 as savourrpc_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x17savourrpc/airdrop.proto\x12\x11savourrpc.airdrop\x1a\x16savourrpc/common.proto\"Y\n\x10\x44ppLinkPointsReq\x12\x16\n\x0e\x63onsumer_token\x18\x01 \x01(\t\x12\x0c\n\x04type\x18\x02 \x01(\x05\x12\x0f\n\x07\x61\x64\x64ress\x18\x03 \x01(\t\x12\x0e\n\x06points\x18\x04 \x01(\r\"D\n\x10\x44ppLinkPointsRep\x12#\n\x04\x63ode\x18\x01 \x01(\x0e\x32\x15.savourrpc.ReturnCode\x12\x0b\n\x03msg\x18\x02 \x01(\t2s\n\x0e\x41irdropService\x12\x61\n\x13submitDppLinkPoints\x12#.savourrpc.airdrop.DppLinkPointsReq\x1a#.savourrpc.airdrop.DppLinkPointsRep\"\x00\x42\'\n\x14group.savour.airdropZ\x0f./proto/airdropb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'savourrpc.airdrop_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\024group.savour.airdropZ\017./proto/airdrop'
  _globals['_DPPLINKPOINTSREQ']._serialized_start=70
  _globals['_DPPLINKPOINTSREQ']._serialized_end=159
  _globals['_DPPLINKPOINTSREP']._serialized_start=161
  _globals['_DPPLINKPOINTSREP']._serialized_end=229
  _globals['_AIRDROPSERVICE']._serialized_start=231
  _globals['_AIRDROPSERVICE']._serialized_end=346
# @@protoc_insertion_point(module_scope)