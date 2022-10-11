import { useEffect, useRef } from "react"

import { nameToColor } from "../common/utils"

interface Audio {
  audio: MediaStream
  name: string
}

export default function AudioVisualizer(props: Audio) {
  const analyserCanvas = useRef(null)
  const color = nameToColor(props.name)
  useEffect(() => {
    if (props.audio === undefined) return
    const audioCtx = new AudioContext()
    const analyser = audioCtx.createAnalyser()
    const audioSrc = audioCtx.createMediaStreamSource(props.audio)
    const distortion = audioCtx.createWaveShaper()
    audioSrc.connect(analyser)
    analyser.connect(distortion)
    distortion.connect(audioCtx.destination)
    analyser.fftSize = 256
    const bufferLength = analyser.frequencyBinCount
    const dataArray = new Uint8Array(bufferLength)
    analyser.getByteTimeDomainData(dataArray)

    const canvasCtx = analyserCanvas.current.getContext("2d")

    const draw = () => {
      requestAnimationFrame(draw)
      analyser.getByteFrequencyData(dataArray)
      const barWidth = 4
      let barHeight: number
      let x = 0

      for (let i = 0; i < bufferLength; i++) {
        barHeight = dataArray[i] / 2

        canvasCtx.fillStyle = `rgb(${barHeight}, ${color[1]}, ${color[2]})`
        canvasCtx.fillRect(x, 20 - barHeight / 2, barWidth, barHeight)

        x += barWidth + 1
      }
    }
    draw()
  }, [props.audio, color])

  return (
    <div className="visualizer mx-auto mt-4">
      <canvas ref={analyserCanvas} className="h-1/2 w-4/5"></canvas>
    </div>
  )
}
