# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: filesystem/filesystem.proto
# Protobuf Python Version: 5.29.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder

# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(
    b'\n\x1b\x66ilesystem/filesystem.proto\x12\nfilesystem"G\n\x0bMoveRequest\x12\x16\n\x06source\x18\x01 \x01(\tR\x06source\x12 \n\x0b\x64\x65stination\x18\x02 \x01(\tR\x0b\x64\x65stination";\n\x0cMoveResponse\x12+\n\x05\x65ntry\x18\x01 \x01(\x0b\x32\x15.filesystem.EntryInfoR\x05\x65ntry"$\n\x0eMakeDirRequest\x12\x12\n\x04path\x18\x01 \x01(\tR\x04path">\n\x0fMakeDirResponse\x12+\n\x05\x65ntry\x18\x01 \x01(\x0b\x32\x15.filesystem.EntryInfoR\x05\x65ntry"#\n\rRemoveRequest\x12\x12\n\x04path\x18\x01 \x01(\tR\x04path"\x10\n\x0eRemoveResponse"!\n\x0bStatRequest\x12\x12\n\x04path\x18\x01 \x01(\tR\x04path";\n\x0cStatResponse\x12+\n\x05\x65ntry\x18\x01 \x01(\x0b\x32\x15.filesystem.EntryInfoR\x05\x65ntry"]\n\tEntryInfo\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12(\n\x04type\x18\x02 \x01(\x0e\x32\x14.filesystem.FileTypeR\x04type\x12\x12\n\x04path\x18\x03 \x01(\tR\x04path"$\n\x0eListDirRequest\x12\x12\n\x04path\x18\x01 \x01(\tR\x04path"B\n\x0fListDirResponse\x12/\n\x07\x65ntries\x18\x01 \x03(\x0b\x32\x15.filesystem.EntryInfoR\x07\x65ntries"C\n\x0fWatchDirRequest\x12\x12\n\x04path\x18\x01 \x01(\tR\x04path\x12\x1c\n\trecursive\x18\x02 \x01(\x08R\trecursive"P\n\x0f\x46ilesystemEvent\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12)\n\x04type\x18\x02 \x01(\x0e\x32\x15.filesystem.EventTypeR\x04type"\xfe\x01\n\x10WatchDirResponse\x12?\n\x05start\x18\x01 \x01(\x0b\x32\'.filesystem.WatchDirResponse.StartEventH\x00R\x05start\x12=\n\nfilesystem\x18\x02 \x01(\x0b\x32\x1b.filesystem.FilesystemEventH\x00R\nfilesystem\x12\x46\n\tkeepalive\x18\x03 \x01(\x0b\x32&.filesystem.WatchDirResponse.KeepAliveH\x00R\tkeepalive\x1a\x0c\n\nStartEvent\x1a\x0b\n\tKeepAliveB\x07\n\x05\x65vent"H\n\x14\x43reateWatcherRequest\x12\x12\n\x04path\x18\x01 \x01(\tR\x04path\x12\x1c\n\trecursive\x18\x02 \x01(\x08R\trecursive"6\n\x15\x43reateWatcherResponse\x12\x1d\n\nwatcher_id\x18\x01 \x01(\tR\twatcherId"8\n\x17GetWatcherEventsRequest\x12\x1d\n\nwatcher_id\x18\x01 \x01(\tR\twatcherId"O\n\x18GetWatcherEventsResponse\x12\x33\n\x06\x65vents\x18\x01 \x03(\x0b\x32\x1b.filesystem.FilesystemEventR\x06\x65vents"5\n\x14RemoveWatcherRequest\x12\x1d\n\nwatcher_id\x18\x01 \x01(\tR\twatcherId"\x17\n\x15RemoveWatcherResponse*R\n\x08\x46ileType\x12\x19\n\x15\x46ILE_TYPE_UNSPECIFIED\x10\x00\x12\x12\n\x0e\x46ILE_TYPE_FILE\x10\x01\x12\x17\n\x13\x46ILE_TYPE_DIRECTORY\x10\x02*\x98\x01\n\tEventType\x12\x1a\n\x16\x45VENT_TYPE_UNSPECIFIED\x10\x00\x12\x15\n\x11\x45VENT_TYPE_CREATE\x10\x01\x12\x14\n\x10\x45VENT_TYPE_WRITE\x10\x02\x12\x15\n\x11\x45VENT_TYPE_REMOVE\x10\x03\x12\x15\n\x11\x45VENT_TYPE_RENAME\x10\x04\x12\x14\n\x10\x45VENT_TYPE_CHMOD\x10\x05\x32\x9f\x05\n\nFilesystem\x12\x39\n\x04Stat\x12\x17.filesystem.StatRequest\x1a\x18.filesystem.StatResponse\x12\x42\n\x07MakeDir\x12\x1a.filesystem.MakeDirRequest\x1a\x1b.filesystem.MakeDirResponse\x12\x39\n\x04Move\x12\x17.filesystem.MoveRequest\x1a\x18.filesystem.MoveResponse\x12\x42\n\x07ListDir\x12\x1a.filesystem.ListDirRequest\x1a\x1b.filesystem.ListDirResponse\x12?\n\x06Remove\x12\x19.filesystem.RemoveRequest\x1a\x1a.filesystem.RemoveResponse\x12G\n\x08WatchDir\x12\x1b.filesystem.WatchDirRequest\x1a\x1c.filesystem.WatchDirResponse0\x01\x12T\n\rCreateWatcher\x12 .filesystem.CreateWatcherRequest\x1a!.filesystem.CreateWatcherResponse\x12]\n\x10GetWatcherEvents\x12#.filesystem.GetWatcherEventsRequest\x1a$.filesystem.GetWatcherEventsResponse\x12T\n\rRemoveWatcher\x12 .filesystem.RemoveWatcherRequest\x1a!.filesystem.RemoveWatcherResponseBi\n\x0e\x63om.filesystemB\x0f\x46ilesystemProtoP\x01\xa2\x02\x03\x46XX\xaa\x02\nFilesystem\xca\x02\nFilesystem\xe2\x02\x16\x46ilesystem\\GPBMetadata\xea\x02\nFilesystemb\x06proto3'
)

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(
    DESCRIPTOR, "filesystem.filesystem_pb2", _globals
)
if not _descriptor._USE_C_DESCRIPTORS:
    _globals["DESCRIPTOR"]._loaded_options = None
    _globals[
        "DESCRIPTOR"
    ]._serialized_options = b"\n\016com.filesystemB\017FilesystemProtoP\001\242\002\003FXX\252\002\nFilesystem\312\002\nFilesystem\342\002\026Filesystem\\GPBMetadata\352\002\nFilesystem"
    _globals["_FILETYPE"]._serialized_start = 1388
    _globals["_FILETYPE"]._serialized_end = 1470
    _globals["_EVENTTYPE"]._serialized_start = 1473
    _globals["_EVENTTYPE"]._serialized_end = 1625
    _globals["_MOVEREQUEST"]._serialized_start = 43
    _globals["_MOVEREQUEST"]._serialized_end = 114
    _globals["_MOVERESPONSE"]._serialized_start = 116
    _globals["_MOVERESPONSE"]._serialized_end = 175
    _globals["_MAKEDIRREQUEST"]._serialized_start = 177
    _globals["_MAKEDIRREQUEST"]._serialized_end = 213
    _globals["_MAKEDIRRESPONSE"]._serialized_start = 215
    _globals["_MAKEDIRRESPONSE"]._serialized_end = 277
    _globals["_REMOVEREQUEST"]._serialized_start = 279
    _globals["_REMOVEREQUEST"]._serialized_end = 314
    _globals["_REMOVERESPONSE"]._serialized_start = 316
    _globals["_REMOVERESPONSE"]._serialized_end = 332
    _globals["_STATREQUEST"]._serialized_start = 334
    _globals["_STATREQUEST"]._serialized_end = 367
    _globals["_STATRESPONSE"]._serialized_start = 369
    _globals["_STATRESPONSE"]._serialized_end = 428
    _globals["_ENTRYINFO"]._serialized_start = 430
    _globals["_ENTRYINFO"]._serialized_end = 523
    _globals["_LISTDIRREQUEST"]._serialized_start = 525
    _globals["_LISTDIRREQUEST"]._serialized_end = 561
    _globals["_LISTDIRRESPONSE"]._serialized_start = 563
    _globals["_LISTDIRRESPONSE"]._serialized_end = 629
    _globals["_WATCHDIRREQUEST"]._serialized_start = 631
    _globals["_WATCHDIRREQUEST"]._serialized_end = 698
    _globals["_FILESYSTEMEVENT"]._serialized_start = 700
    _globals["_FILESYSTEMEVENT"]._serialized_end = 780
    _globals["_WATCHDIRRESPONSE"]._serialized_start = 783
    _globals["_WATCHDIRRESPONSE"]._serialized_end = 1037
    _globals["_WATCHDIRRESPONSE_STARTEVENT"]._serialized_start = 1003
    _globals["_WATCHDIRRESPONSE_STARTEVENT"]._serialized_end = 1015
    _globals["_WATCHDIRRESPONSE_KEEPALIVE"]._serialized_start = 1017
    _globals["_WATCHDIRRESPONSE_KEEPALIVE"]._serialized_end = 1028
    _globals["_CREATEWATCHERREQUEST"]._serialized_start = 1039
    _globals["_CREATEWATCHERREQUEST"]._serialized_end = 1111
    _globals["_CREATEWATCHERRESPONSE"]._serialized_start = 1113
    _globals["_CREATEWATCHERRESPONSE"]._serialized_end = 1167
    _globals["_GETWATCHEREVENTSREQUEST"]._serialized_start = 1169
    _globals["_GETWATCHEREVENTSREQUEST"]._serialized_end = 1225
    _globals["_GETWATCHEREVENTSRESPONSE"]._serialized_start = 1227
    _globals["_GETWATCHEREVENTSRESPONSE"]._serialized_end = 1306
    _globals["_REMOVEWATCHERREQUEST"]._serialized_start = 1308
    _globals["_REMOVEWATCHERREQUEST"]._serialized_end = 1361
    _globals["_REMOVEWATCHERRESPONSE"]._serialized_start = 1363
    _globals["_REMOVEWATCHERRESPONSE"]._serialized_end = 1386
    _globals["_FILESYSTEM"]._serialized_start = 1628
    _globals["_FILESYSTEM"]._serialized_end = 2299
# @@protoc_insertion_point(module_scope)
