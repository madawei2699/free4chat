// pure components and storybook-like sandbox
import React, { useState } from "react";
import css from "./VoiceChat.module.css";
import { User } from "./api";
import { useAudioContext, AudioContextProvider } from "./VoiceChat";

export const Storybook: React.FC = () => {
  return (
    <div style={{ padding: 10 }}>
      <div>
        <h3>icons</h3>
        <ButtonMicrohoneContainer />
        <ButtonSpeakerContainer />

        <div style={{ width: 400 }}>
          <h3>user me</h3>
          <UserMe
            user={{
              emoji: "😎",
              id: "123",
              mute: false,
            }}
            isMutedMicrophone={true}
            isMutedSpeaker={false}
            onClickMuteMicrohone={() => undefined}
            onClickMuteSpeaker={() => undefined}
          />
        </div>
        <div style={{ width: 400 }}>
          <h3>user remote</h3>
          <UserRemoteStory />
        </div>
        <div style={{ width: 400 }}>
          <h3>users list</h3>
          <UsersRemoteListStory />
        </div>
        <div>
          <h3>empty room</h3>
          <EmptyRoomStory />
        </div>
        <div>
          <h3>microphone volume analyser</h3>
          <MicrophoneVolumeAnalyser />
        </div>
      </div>
    </div>
  );
};

const getMediaStreamVolume = (
  audioContext: AudioContext,
  mediaStream: MediaStream,
  callback: (isSpeaking: boolean) => void
) => {
  const analyser = audioContext.createAnalyser();
  const mediaStreamSource = audioContext.createMediaStreamSource(mediaStream);
  const processor = audioContext.createScriptProcessor(512);

  analyser.smoothingTimeConstant = 0.8;
  analyser.fftSize = 1024;

  mediaStreamSource.connect(analyser);
  analyser.connect(processor);
  processor.connect(audioContext.destination);
  processor.onaudioprocess = function (event) {
    const buf = event.inputBuffer.getChannelData(0);
    let sum = 0;
    for (let i = 0; i < buf.length; i++) {
      const x = buf[i];
      sum += x * x;
    }
    const rms = Math.sqrt(sum / buf.length);
    if (rms > 0.05) {
      callback(true);
    } else {
      callback(false);
    }
    // callback(rms);
  };
};

const MicrophoneVolumeAnalyserWrapper = () => {
  const audioContext = useAudioContext();
  return (
    <div>
      <button
        onClick={async () => {
          try {
            if (audioContext.state === "suspended") {
              alert("suspended");
              await audioContext.resume();
            }
            const stream = await navigator.mediaDevices.getUserMedia({
              audio: true,
              video: false,
            });
            console.log(stream);
            getMediaStreamVolume(audioContext, stream, (volume) => {
              console.log(volume);
            });
          } catch (error) {
            console.error(error);
          }
        }}
      >
        request
      </button>
    </div>
  );
};
const MicrophoneVolumeAnalyser = () => {
  return (
    <AudioContextProvider>
      <MicrophoneVolumeAnalyserWrapper />
    </AudioContextProvider>
  );
};

const EmptyRoomStory = () => {
  return (
    <div
      style={{ backgroundColor: "rgba(0,0,0,0.1)", height: 400, width: 400 }}
    >
      <EmptyRoom />
    </div>
  );
};
export const EmptyRoom = () => {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        height: "100%",
        justifyContent: "center",
      }}
    >
      <div
        style={{
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          fontSize: 60,
        }}
      >
        👀
      </div>
      <div style={{ textAlign: "center", fontSize: 16, color: "white" }}>
        Waiting for the others...
      </div>
    </div>
  );
};

const UsersRemoteListStory = () => {
  return (
    <div style={{ backgroundColor: "rgba(0,0,0,0.1)" }}>
      <UsersRemoteList
        users={[
          { emoji: "😎", id: "123", mute: false },
          { emoji: "😎", id: "123", mute: true },
          { emoji: "😎", id: "123", mute: false },
        ]}
      />
    </div>
  );
};
export const UsersRemoteList: React.FC<{
  users: User[];
}> = ({ users }) => {
  const renderUsers = () => {
    return users.map((user) => {
      return (
        <div style={{ margin: 5 }} key={user.id}>
          <UserRemote user={user} />
        </div>
      );
    });
  };
  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      {renderUsers()}
    </div>
  );
};

export const UserMe: React.FC<{
  user: User;
  isMutedMicrophone: boolean;
  isMutedSpeaker: boolean;
  onClickMuteMicrohone: (event: React.MouseEvent) => void;
  onClickMuteSpeaker: (event: React.MouseEvent) => void;
}> = ({
  user,
  isMutedMicrophone,
  isMutedSpeaker,
  onClickMuteMicrohone,
  onClickMuteSpeaker,
}) => {
  return (
    <div>
      <div
        style={{
          display: "flex",
          justifyContent: "center",
          pointerEvents: "none",
          userSelect: "none",
        }}
      >
        <div style={{ fontSize: 96 }}>{user.emoji}</div>
      </div>
      <div className={css.buttons}>
        <ButtonMicrohone
          muted={isMutedMicrophone}
          onClick={(event) => {
            onClickMuteMicrohone(event);
          }}
        />
        <ButtonSpeaker
          muted={isMutedSpeaker}
          onClick={(event) => {
            onClickMuteSpeaker(event);
          }}
        />
      </div>
    </div>
  );
};

const UserRemoteStory = () => {
  const [isMuted, setIsMuted] = useState(true);
  return (
    <div>
      <div>
        <label>
          <input
            type="checkbox"
            onChange={() => setIsMuted(!isMuted)}
            checked={isMuted}
          />
          muted
        </label>
      </div>
      <div style={{ backgroundColor: "rgba(0,0,0,0.1)" }}>
        <UserRemote user={{ emoji: "😎", id: "1", mute: isMuted }} />
      </div>
    </div>
  );
};
export const UserRemote: React.FC<{
  user: User;
}> = ({ user }) => {
  return (
    <div
      style={{
        width: 60,
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        pointerEvents: "none",
        userSelect: "none",
        position: "relative",
        backgroundColor: "rgba(0,0,0,0.1)",
        borderRadius: 16,
        padding: 10,
      }}
      key={user.id}
    >
      <div style={{ fontSize: 48 }}>
        <span>{user.emoji}</span>
      </div>
      {user.mute && (
        <div
          style={{
            height: 20,
            width: 20,
            position: "absolute",
            bottom: 5,
            left: 5,
          }}
        >
          <IconMicrophone muted />
        </div>
      )}
    </div>
  );
};

const ButtonMicrohoneContainer = () => {
  const [isMuted, setIsMuted] = useState<boolean>(false);
  return (
    <ButtonMicrohone muted={isMuted} onClick={() => setIsMuted(!isMuted)} />
  );
};

const ButtonSpeakerContainer = () => {
  const [isMuted, setIsMuted] = useState<boolean>(false);
  return <ButtonSpeaker muted={isMuted} onClick={() => setIsMuted(!isMuted)} />;
};

export const ButtonMicrohone: React.FC<{
  muted: boolean;
  onClick: (event: React.MouseEvent) => void;
}> = ({ onClick, muted }) => {
  return (
    <button className={css.buttonMicrophone} onClick={onClick}>
      <IconMicrophone muted={muted} />
    </button>
  );
};

export const ButtonSpeaker: React.FC<{
  muted: boolean;
  onClick: (event: React.MouseEvent) => void;
}> = ({ onClick, muted }) => {
  return (
    <button className={css.buttonMicrophone} onClick={onClick}>
      <IconSpeaker muted={muted} />
    </button>
  );
};

const IconMicrophone: React.FC<{
  muted: boolean;
}> = ({ muted }) => {
  if (muted) {
    return (
      <svg aria-hidden="false" width="100%" height="100%" viewBox="0 0 24 24">
        <path
          d="M6.7 11H5C5 12.19 5.34 13.3 5.9 14.28L7.13 13.05C6.86 12.43 6.7 11.74 6.7 11Z"
          fill="white"
        ></path>
        <path
          d="M9.01 11.085C9.015 11.1125 9.02 11.14 9.02 11.17L15 5.18V5C15 3.34 13.66 2 12 2C10.34 2 9 3.34 9 5V11C9 11.03 9.005 11.0575 9.01 11.085Z"
          fill="white"
        ></path>
        <path
          d="M11.7237 16.0927L10.9632 16.8531L10.2533 17.5688C10.4978 17.633 10.747 17.6839 11 17.72V22H13V17.72C16.28 17.23 19 14.41 19 11H17.3C17.3 14 14.76 16.1 12 16.1C11.9076 16.1 11.8155 16.0975 11.7237 16.0927Z"
          fill="white"
        ></path>
        <path
          d="M21 4.27L19.73 3L3 19.73L4.27 21L8.46 16.82L9.69 15.58L11.35 13.92L14.99 10.28L21 4.27Z"
          fill="#f04747"
        ></path>
      </svg>
    );
  }

  return (
    <svg aria-hidden="false" width="100%" height="100%" viewBox="0 0 24 24">
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M14.99 11C14.99 12.66 13.66 14 12 14C10.34 14 9 12.66 9 11V5C9 3.34 10.34 2 12 2C13.66 2 15 3.34 15 5L14.99 11ZM12 16.1C14.76 16.1 17.3 14 17.3 11H19C19 14.42 16.28 17.24 13 17.72V21H11V17.72C7.72 17.23 5 14.41 5 11H6.7C6.7 14 9.24 16.1 12 16.1ZM12 4C11.2 4 11 4.66667 11 5V11C11 11.3333 11.2 12 12 12C12.8 12 13 11.3333 13 11V5C13 4.66667 12.8 4 12 4Z"
        fill="white"
      ></path>
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M14.99 11C14.99 12.66 13.66 14 12 14C10.34 14 9 12.66 9 11V5C9 3.34 10.34 2 12 2C13.66 2 15 3.34 15 5L14.99 11ZM12 16.1C14.76 16.1 17.3 14 17.3 11H19C19 14.42 16.28 17.24 13 17.72V22H11V17.72C7.72 17.23 5 14.41 5 11H6.7C6.7 14 9.24 16.1 12 16.1Z"
        fill="white"
      ></path>
    </svg>
  );
};

const IconSpeaker: React.FC<{
  muted: boolean;
}> = ({ muted }) => {
  if (muted) {
    return (
      <svg aria-hidden="false" width="100%" height="100%" viewBox="0 0 24 24">
        <path
          d="M6.16204 15.0065C6.10859 15.0022 6.05455 15 6 15H4V12C4 7.588 7.589 4 12 4C13.4809 4 14.8691 4.40439 16.0599 5.10859L17.5102 3.65835C15.9292 2.61064 14.0346 2 12 2C6.486 2 2 6.485 2 12V19.1685L6.16204 15.0065Z"
          fill="white"
        ></path>
        <path
          d="M19.725 9.91686C19.9043 10.5813 20 11.2796 20 12V15H18C16.896 15 16 15.896 16 17V20C16 21.104 16.896 22 18 22H20C21.105 22 22 21.104 22 20V12C22 10.7075 21.7536 9.47149 21.3053 8.33658L19.725 9.91686Z"
          fill="white"
        ></path>
        {muted && (
          <path
            d="M3.20101 23.6243L1.7868 22.2101L21.5858 2.41113L23 3.82535L3.20101 23.6243Z"
            fill="#f04747"
          ></path>
        )}
      </svg>
    );
  }

  return (
    <svg aria-hidden="false" width="100%" height="100%" viewBox="0 0 24 24">
      <svg width="24" height="24" viewBox="0 0 24 24">
        <path
          d="M12 2.00305C6.486 2.00305 2 6.48805 2 12.0031V20.0031C2 21.1071 2.895 22.0031 4 22.0031H6C7.104 22.0031 8 21.1071 8 20.0031V17.0031C8 15.8991 7.104 15.0031 6 15.0031H4V12.0031C4 7.59105 7.589 4.00305 12 4.00305C16.411 4.00305 20 7.59105 20 12.0031V15.0031H18C16.896 15.0031 16 15.8991 16 17.0031V20.0031C16 21.1071 16.896 22.0031 18 22.0031H20C21.104 22.0031 22 21.1071 22 20.0031V12.0031C22 6.48805 17.514 2.00305 12 2.00305Z"
          fill="white"
        ></path>
      </svg>
    </svg>
  );
};
