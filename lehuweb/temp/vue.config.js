const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      "/api": {
        //访问时'/api'会被自动替换为target中的网址
        target: "http://127.0.0.1:3000/api", // 需要跨域访问的网址
        changeOrigin: true,
        pathRewrite: {
          "^/api": "",
        },
      },
    },
  },
});
