import React from "react";
import { ImageBackground, StyleSheet, Text, View, SafeAreaView, Image, Alert, TouchableOpacity } from "react-native";

const App = () => (
  <SafeAreaView style={styles.container}>
    <View>
      <ImageBackground 
        source={require('./assets/graph-texture.jpeg')} 
        resizeMode="cover" 
        style={styles.imageLogo}
      >
        <View style={styles.upper}>
          <View style={styles.logo}>
            <Text style={styles.textLogo}>
              retelling
            </Text>
          </View>
          <TouchableOpacity
            onPress={()=>alert('Ошибочка вышла')}
            style={styles.accButton}
          >
            <Text>Sign Up</Text>
          </TouchableOpacity>
        </View>
      </ImageBackground>
    </View>
    
    <View>
      <TouchableOpacity
        onPress={() => alert('Ничего не найдено!')}
        style={styles.defButton}
      >
        <View style={{flexDirection:'row', alignItems:'left'}}>
        <Image 
          source={require('./assets/premium-icon-book-828370.png')}
          style={styles.btnIcon}  
        />
        <Text style={styles.defButtonText}>Книги</Text>
        </View>
      </TouchableOpacity>
      <TouchableOpacity
        onPress={() => alert('Ничего не найдено!')}
        style={styles.defButton}
      >
        <View style={{flexDirection:'row', alignItems:'left'}}>
        <Image 
          source={require('./assets/premium-icon-book-828370.png')}
          style={styles.btnIcon}  
        />
        <Text style={styles.defButtonText}>Фильмы</Text>
        </View>
      </TouchableOpacity>
      <TouchableOpacity
        onPress={() => alert('Ничего не найдено!')}
        style={styles.defButton}
      >
        <View style={{flexDirection:'row', alignItems:'left'}}>
        <Image 
          source={require('./assets/premium-icon-book-828370.png')}
          style={styles.btnIcon}  
        />
        <Text style={styles.defButtonText}>Видеоигры</Text>
        </View>
      </TouchableOpacity>
    </View>
  </SafeAreaView>
);

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  imageLogo: {
    height: 84,
  },
  logo: {
  },
  textLogo: {
    fontFamily: "Hiragino Sans",
    fontSize: 42,
    lineHeight: 84,
    color: "white",
    textAlign: "left",
    marginLeft: 20,
  },
  accButton: {
    elevation: 8,
    borderColor: "#000",
    backgroundColor: "#fff",
    borderWidth: 2,
    borderRadius: 10,
    paddingVertical: 10,
    paddingHorizontal: 12,
    margin: 10,
  },
  defButton: {
    elevation: 8,
    backgroundColor: "#fff",
    borderColor: "#000",
    borderWidth: 2,
    borderRadius: 10,
    paddingVertical: 10,
    paddingHorizontal: 12,
    margin: 10,
  },
  defButtonText: {
    fontSize: 15,
  },
  btnIcon: {
    width:15,
    height:15,
    margin: 2,
    marginRight: 5,
  },
  upper: {
    flex:1,flexDirection:"row",justifyContent:"space-between",backgroundColor: "#000000c0",}
  
});

export default App;