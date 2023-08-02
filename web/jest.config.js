module.exports = {
    testEnvironment: 'jsdom',
    testMatch:[
      "**/?(*.)+(spec|test|unit).[jt]s?(x)" // 匹配单元测试文件
    ],
    transform: {
        "^.+\.[j|t]sx?$": "babel-jest",
        "^.+\.vue?$": "@vue/vue3-jest",
        "^.+\.tsx$": "ts-jest"
    }
}
