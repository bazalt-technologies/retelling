import createPersistedState from 'vuex-persistedstate'
const storeData = {
    state() {
        return {
            user: null,
            content: [],
            types: [],
            genres: [],
            allUsers: [],
            selectedContent: null
        }
    },
    getters: {
        getUser(state) {
          return state.user
        },
        getTypes(state) {
          return state.types
        },
        getGenres(state) {
          return state.genres
        },
        getContent(state) {
          return state.content
        },
        getAllUsers(state) {
            return state.allUsers
        },
        getSelectedContent(state) {
            return state.selectedContent
        }
    },
    plugins: [createPersistedState()],
    mutations: {
        setUser(state, user) {
            state.user = user
        },
        setContent(state, content) {
            state.content = content
        },
        setGenres(state, genres) {
            state.genres = genres
        },
        setTypes(state, types) {
            state.types = types
        },
        setAllUsers(state, users) {
            state.allUsers = users
        },
        setSelectedContent(state, selectedContent) {
            state.selectedContent = selectedContent
        }
    }
}

export default storeData