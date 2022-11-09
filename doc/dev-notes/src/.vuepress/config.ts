import { defineUserConfig } from "vuepress";
import theme from "./theme.js";

export default defineUserConfig({
  base: "/",

  locales: {
    "/": {
      lang: "zh-CN",
      title: "free4chat",
      description: "使用Elixir开发实时Web应用",
    },
  },

  theme,

  shouldPrefetch: false,
});
