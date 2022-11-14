import { sidebar } from "vuepress-theme-hope";

export const zhSidebar = sidebar({
  "/": [
    {
      text: "开发手记",
      icon: "screwdriver-wrench",
      prefix: "note/",
      children: "structure",
    },
    {
      text: "演讲稿",
      icon: "video",
      prefix: "slide/",
      children: "structure",
    },
  ],
});
