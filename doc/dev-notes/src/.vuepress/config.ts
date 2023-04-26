import { googleAnalyticsPlugin } from "@vuepress/plugin-google-analytics";
import { searchPlugin } from '@vuepress/plugin-search'
import { defineUserConfig } from "vuepress";
import theme from "./theme.js";

export default defineUserConfig({
  base: "/",

  head: [
    ['script', {src: 'https://umami.bmpi.dev/script.js', 'data-website-id': '756d3289-e627-409c-a9d7-a272cf789f0b', async: true, defer: true}]
  ],

  locales: {
    "/": {
      lang: "zh-CN",
      title: "free4chat",
      description: "使用Elixir开发实时Web应用",
    },
  },

  theme,

  shouldPrefetch: false,

  plugins: [
      googleAnalyticsPlugin({
        id: "G-8HBGR25L32",
      }),
      searchPlugin({  
        locales: {
          "/": {
            placeholder: "搜索文档",
          },
        },
      }),
    ]
});
