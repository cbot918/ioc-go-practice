class Cheer {
  Hello() {
    throw new Error("method1 must be implemented");
  }

  HowAreYou() {
    throw new Error("method2 must be implemented");
  }
}

class Chinese extends Cheer {
  Hello() {
    console.log("安安");
  }

  HowAreYou() {
    console.log("你好嗎");
  }
}

class English extends Cheer {
  Hello() {
    console.log("hello");
  }

  HowAreYou() {
    console.log("how are you");
  }
}

function main() {
  run(new Chinese());
  run(new English());
}
main();

function run(service) {
  service.Hello();
  service.HowAreYou();
}
