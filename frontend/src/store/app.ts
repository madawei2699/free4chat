export const state = () => ({
  snackbar: {
    show: false,
    message: '',
    color: 'info'
  },
  chat: {
    rooms: {},
    mutes: {}
  },
  profile: {
    uid: ''
  },
  appbar: {
    title: '',
    style: '',
    show: true,
    back: false,
    home: false,
    dark: true,
    animation: false,
    color: '#111'
  }
})

export const mutations = {
  toast (state, { message, color }) {
    state.snackbar.show = true
    state.snackbar.message = message
    state.snackbar.color = color
  },
  setSnackbar (state, val) {
    state.snackbar.show = val
  },
  SET_PROFILE (state, val) {
    state.profile = val
  },
  SET_APPBAR (state, value) {
    const defaultValue = {
      title: '',
      style: '',
      show: true,
      back: false,
      home: false,
      animation: false,
      dark: true,
      color: '#111'
    }
    state.appbar = { ...defaultValue, ...value }
  },
  ADD_ROOM (state, { room, nickname }) {
    if (!state.chat.rooms.hasOwnProperty(room)) {
      state.chat.rooms[room] = { room, nickname }
    }
  },
  REMOVE_ROOM (state, { room }) {
    if (state.chat.rooms.hasOwnProperty(room)) {
      const rms = state.chat.rooms
      delete rms[room]
      state.chat.rooms = Object.assign({}, rms)
    }
  },
  ADD_MUTE (state, uid) {
    const mutes = state.chat.mutes
    if (!state.chat.mutes.hasOwnProperty(uid)) {
      mutes[uid] = 1
    }
    state.chat.mutes = Object.assign({}, mutes)
  },
  REMOVE_MUTE (state, uid) {
    if (state.chat.mutes.hasOwnProperty(uid)) {
      const mutes = state.chat.mutes
      delete mutes[uid]
      state.chat.mutes = Object.assign({}, mutes)
    }
  },
  ADD_MUTES (state, uids) {
    const mutes = state.chat.mutes
    for (let ix = 0; ix < uids.length; ix++) {
      if (!mutes.hasOwnProperty(uids[ix])) {
        mutes[uids[ix]] = 1
      }
    }
    state.chat.mutes = Object.assign({}, mutes)
  },
  REMOVE_MUTES (state, uids) {
    const mutes = state.chat.mutes
    for (let ix = 0; ix < uids.length; ix++) {
      if (state.chat.mutes.hasOwnProperty(uids[ix])) {
        delete mutes[uids[ix]]
      }
    }
    state.chat.mutes = Object.assign({}, mutes)
  },
  SET_NICKNAME (state, { room, nickname }) {
    state.chat.rooms[room] = { room, nickname }
  }
}
