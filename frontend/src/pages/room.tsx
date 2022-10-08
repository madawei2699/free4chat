import React, { useState, useEffect } from "react"

import Head from "next/head"
import { useRouter } from "next/router"

import { randomName, saveRoomToLocalStorage } from "../common/utils"

export default function Room() {
  const router = useRouter()
  const roomId = router.query.id as string
  const [roomName, setRoomName] = useState<string>("")
  const [nickName, setNickName] = useState<string>("")
  const [showNickNamePop, setShowNickNamePop] = useState<boolean>(false)

  const dissmisNickNamePop = () => {
    setShowNickNamePop(false)
    saveRoomToLocalStorage(roomName, nickName)
  }

  useEffect(() => {
    if (!router.isReady) return
    if (roomId === undefined || roomId === "") {
      router.push("/")
    }
    setRoomName(roomId)
    if (typeof window !== "undefined") {
      var rooms: {}[] = JSON.parse(localStorage.getItem("rooms") || "[]")
      const room: any = rooms.find((room: any) => room.roomName === roomId)
      if (room !== undefined) {
        setNickName(room.nickName)
      } else {
        setShowNickNamePop(true)
      }
    }
  }, [setRoomName, roomId, router])

  return (
    <div>
      <Head>
        <title>Room#{roomName} - Free4Chat</title>
      </Head>

      <main className="bg-gray-900 text-white">
        {(showNickNamePop || (roomName !== "" && nickName === "")) && (
          <div className="m-auto max-w-md rounded-lg bg-gray-800 p-8 shadow-2xl">
            <h2 className="text-lg font-bold">Input Your Nick Name</h2>
            <div className="relative mb-4 mt-4">
              <label htmlFor="nickname" className="sr-only">
                Nickname
              </label>

              <input
                type="input"
                id="nickname"
                value={nickName}
                placeholder="Nick Name"
                onChange={(e) => setNickName(e.target.value)}
                className="w-full rounded-md border-gray-200 bg-white p-3 text-gray-700 shadow-sm transition focus:border-white focus:outline-none focus:ring focus:ring-yellow-400"
              />

              <span className="absolute inset-y-0 right-0 grid w-10 place-content-center">
                <button
                  onClick={() => setNickName(randomName())}
                  type="button"
                  className="text-black hover:bg-red-400"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 16 16"
                    fill="currentColor"
                    className="h-4 w-4"
                  >
                    <path d="M13 1a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2h10zM3 0a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V3a3 3 0 0 0-3-3H3z" />{" "}
                    <path d="M5.5 4a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0zm8 0a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0zm0 8a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0zm-8 0a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0zm4-4a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0z" />
                  </svg>
                </button>
              </span>
            </div>

            <button
              type="button"
              onClick={() => dissmisNickNamePop()}
              className="group mt-4 flex w-full items-center justify-center rounded-md bg-rose-600 px-5 py-3 text-white transition focus:outline-none focus:ring focus:ring-yellow-400 sm:mt-0 sm:w-auto"
            >
              <span className="text-sm font-medium"> Go </span>

              <svg
                className="ml-3 h-5 w-5"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M17 8l4 4m0 0l-4 4m4-4H3"
                />
              </svg>
            </button>
          </div>
        )}
        <div className="mx-auto h-screen max-w-screen-xl px-4 py-32 lg:flex lg:items-center"></div>
      </main>
    </div>
  )
}
