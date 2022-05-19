const { expect } = require("chai");

const hw = require("../src/handler/hello_world.js");

const _helloWorld = "Hello world!"

it("Test Hello World", function () {
  res = hw.helloWorld()

  expect(res).to.equal(_helloWorld)
});
