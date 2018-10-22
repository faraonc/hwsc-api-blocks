const HWSC_FILE_METADATA_SVC_PROTO_PATH = __dirname + "/proto/hwsc-file-metadata-svc.proto";
const grpc = require("grpc");
const protoLoader = require("@grpc/proto-loader");


const options = {
    includeDirs: [
        HWSC_FILE_METADATA_SVC_PROTO_PATH
    ]
};
const hwscFileMetadataSvcProtoPkgDef = protoLoader.loadSync("", options);
const hwscFileMetadataSvcPbJs = grpc.loadPackageDefinition(hwscFileMetadataSvcProtoPkgDef).hwscFileMetadataSvc;

function createFileMetadata(fileMetadata, callback) {
    const client = new hwscFileMetadataSvcPbJs.FileMetadataService("localhost:50051",
        grpc.credentials.createInsecure());

    const request = {
        data: fileMetadata
    };

    client.createFileMetadata(request, function (err, response) {
        if (!err) {
            console.log(request);
            grpc.closeClient(client);
        }

        callback(err, response);
    });
}

module.exports = {
    createFileMetadata: createFileMetadata
};
