const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');
const base64 = require('js-base64').Base64;

const PROTO_PATH = `${__dirname}/../..`;

const options = {
  includeDirs: [PROTO_PATH],
};

// runtime dynamic compialation proto file, creates pb file dynamically
const packageDefinition = protoLoader
  .loadSync('protobuf/hwsc-app-gateway-svc/app/hwsc-app-gateway-svc.proto', options);

// pointing to pb file in memory
const protoDescriptor = grpc.loadPackageDefinition(packageDefinition).app;

// insecure no SSL and TLS
const client = new protoDescriptor
  .AppGatewayService('localhost: 50055', grpc.credentials.createInsecure());

const userPass = new grpc.Metadata();
userPass.add('authorization', `Basic ${base64.encode('change-me')}`);

const callbackErr = () => console.error('callback not a function');

function getStatus(callback) {
  if (typeof callback !== 'function') {
    callbackErr();
    return;
  }

  client.getStatus({}, userPass, (err, response) => {
    if (!err) {
      grpc.closeClient(client);
    }
    callback(err, response);
  });
}

module.exports = {
  getStatus,
};
