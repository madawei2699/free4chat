import { sidebar } from "vuepress-theme-hope";

export const zhSidebar = sidebar({
  "/": [
    {
      text: "开发手记",
      icon: "note",
      prefix: "note/",
      children: "structure",
    },
    {
      text: "演讲稿",
      icon: "slides",
      prefix: "slide/",
      children: "structure",
    },
  ],
});
