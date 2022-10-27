const storeData = {
    state() {
        return {
            user: null,
            content: [],
            types: [],
            genres: [],
            allUsers: []
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
        }
    },
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
        }
    }
}

export default storeData