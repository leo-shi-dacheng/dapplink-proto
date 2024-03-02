# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from savourrpc import airdrop_pb2 as savourrpc_dot_airdrop__pb2


class AirdropServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.submitDppLinkPoints = channel.unary_unary(
                '/savourrpc.airdrop.AirdropService/submitDppLinkPoints',
                request_serializer=savourrpc_dot_airdrop__pb2.DppLinkPointsReq.SerializeToString,
                response_deserializer=savourrpc_dot_airdrop__pb2.DppLinkPointsRep.FromString,
                )


class AirdropServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def submitDppLinkPoints(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AirdropServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'submitDppLinkPoints': grpc.unary_unary_rpc_method_handler(
                    servicer.submitDppLinkPoints,
                    request_deserializer=savourrpc_dot_airdrop__pb2.DppLinkPointsReq.FromString,
                    response_serializer=savourrpc_dot_airdrop__pb2.DppLinkPointsRep.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'savourrpc.airdrop.AirdropService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class AirdropService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def submitDppLinkPoints(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/savourrpc.airdrop.AirdropService/submitDppLinkPoints',
            savourrpc_dot_airdrop__pb2.DppLinkPointsReq.SerializeToString,
            savourrpc_dot_airdrop__pb2.DppLinkPointsRep.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)