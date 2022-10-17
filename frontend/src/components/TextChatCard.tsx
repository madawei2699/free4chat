import React, { useState } from "react"

import { LOCAL_PEER_ID } from "@common/consts"
import { Message } from "@common/types"
import { strToBgColor } from "@common/utils"

import Store from "../store/store"

interface Messages {
  room: string
  messages: Message[]
}

export default function TextChatCard(m: Messages) {
  const [message, setMessage] = useState<string>("")
  const sendMessage = (event) => {
    if (event.key === "Enter" && message !== "") {
      Store.sendTextMessage(m.room, message)
      setMessage("")
    }
  }
  return (
    <div className="mt-4 block rounded-xl border border-gray-700 bg-gray-800 p-4 shadow-xl">
      {m.messages.length > 0 && (
        <div className="flex max-h-96 w-full flex-col justify-between overflow-scroll text-sm">
          <div className="mt-5 flex flex-col-reverse">
            {m.messages.map((p, i) =>
              p.peerId === LOCAL_PEER_ID ? (
                <div className="mb-4 flex justify-end" key={i}>
                  <div className="text-white">
                    <div
                      className="mr-2 rounded-bl-3xl rounded-tl-3xl rounded-tr-xl px-4 py-3"
                      style={{
                        backgroundColor: strToBgColor(p.name),
                      }}
                    >
                      {p.text}
                    </div>
                  </div>
                </div>
              ) : (
                <div className="mb-4 flex justify-start" key={i}>
                  <div className="text-white">
                    <div
                      className="ml-2 rounded-br-3xl rounded-tr-3xl rounded-tl-xl px-4 py-3"
                      style={{
                        backgroundColor: strToBgColor(p.name),
                      }}
                    >
                      {p.text}
                    </div>
                  </div>
                </div>
              )
            )}
          </div>
        </div>
      )}

      <div className="mt-2">
        <input
          className="w-full rounded-xl bg-gray-900"
          type="text"
          value={message}
          onKeyDown={sendMessage}
          onChange={(e) => setMessage(e.target.value)}
          placeholder="type your message here..."
        />
      </div>
    </div>
  )
}
