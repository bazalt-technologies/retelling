import requests
from random import randint
import datetime
import time

Names = ["Gleb", "Misha", "Maks", "Anton", "Sasha", "Fedya", "Andrew", "Alex", "Nikita", "Nata", "Olya",
         "Ann", "Georgiy", "Vasiliy", "Steve", "Masha", "Tanya", "Kolya", "Vladimir", "Boris", "Pavel"]

Logins = ["Gleb228", "Misha228", "Maks228", "SuperKiller", "MisterX", "Ovakado", "ToastWithGarlic",
          "Argentuz", "Argentus", "Mishgun", "Agent007", "Pizza", "Peppa", "MoiLOgin1337", "Andrew2003", "Boris86",
          "VladimirVladmirovich", "Vladik2010", "Kashalot", "Mushroom", "DedSGori", "RedmiNote8", "Margarita", "Papiroska",
          "LoginLogin", "Girlyandochka", "MolchatDoma", "Klaviaturochka", "LetiBilOsnovanV1886", "SuperGeniy"]

Professions = ["It-developer", "Cleaner", "Waiter", "Prisoner", "Pirate", "Sailor", "Janitor", "Steve Jobs", "Bill Gates"]


userIDs = []
contentIDs = []

for i in range(40):
    Login = Logins[randint(0, len(Logins) - 1)]
    Name = Names[randint(0, len(Names) - 1)]
    Profession = Professions[randint(0, len(Professions) - 1)]
    response = requests.post("http://127.0.0.1:8081/api/v1/users", json = {"Login" : Login, "Password" : Login,
                                                                           "Data" : {"Name" : Name, "Age" : randint(17, 65), "Profession" : Profession}})
    print(response)
    try:
      userIDs.append(int(response.content.decode()))
    except:
      pass

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID" : 2, "GenreID1" : 1,
                                                                           "GenreID2" : 7, "GenreID3" : 4,
                                                                         "Title" : "Адреналин (2006)",
                                                                         "Description" : "Наемный убийца Чев узнает, что недоброжелатели отравили его редким китайским ядом и отрава начнет действовать немедленно, как только пульс перестанет биться ниже определенной отметки."
                                                                                                                       "\nИ теперь Чеву нужно успеть сделать все свои дела, попробовать раздобыть противоядие и отомстить своим отравителям в прямом смысле слова впопыхах - стараясь создавать вокруг себя как можно более напряженную обстановку и не расслабляться ни на секунду."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))
ContentID = int(response.content.decode())

Date = int(datetime.datetime.now().timestamp())
print(Date)

response = requests.post("http://127.0.0.1:8081/api/v1/reviews", json = {"Date" : Date, "UserID" : userIDs[randint(0, len(userIDs) - 1)], "ContentID" : ContentID, "Review" : "90 минут чистейшего драйва!\nШёл на фильм из-за Дж.Стэтема, ставшего безумно известным после своих ролей в 'Перевозчиках' и культовом 'Большом Куше' не менее культового режиссёра Гая Ритчи. Фильм захватывает с самых первых секунд. Бешеная скорость происходящего и нестандартные ракурсы съёмки захватывают дух и 90 минут пролетают как миг! Фильм просто обалденный, игра актёров изумительна, саундтрек под стать всему фильму такой же агрессивный и резкий. Посмотрите этот фильм, уверен, получите массу удовольствия и новых ощущений, а я же ставлю несомненные.\n10 из 10\nДа здравствует драйв!"})
print(response)
print(response.content.decode())


response = requests.post("http://127.0.0.1:8081/api/v1/reviews", json = {"Date" : Date, "UserID" : userIDs[randint(0, len(userIDs) - 1)], "ContentID" : ContentID, "Review" : "Жизнь - это не те дни, которые прошли,\nа те, которые запомнились.\n\nГабриэль Гарсия Маркес\n=====\n\nОчнувшись, Чев Челиос узнает, что он был убит. Пока он был в «отключке», ему вкололи страшную китайскую отраву – и теперь, если уровень адреналина у него в крови упадет ниже определенной отметки, Чев умрёт.\n\nС этого начинается бешенный и угарный марафон по городу, во время которого Чев успеет разгромить универмаг, полягаться с чёрными байкерами, встретится с боссом, доктором и девушкой и ещё много чего интересного.\n\n«Адреналин», эта темная лошадка, снятая двумя вундеркиндами-дебютантами за 12 миллионов долларов, - больше, чем фильм, и больше, чем развлечение. Это какой-то энергетический коктейль, где агрессия смешана с безудержным весельем, экшен – с философским эликсиром. Впитанный мозгом, «Адреналин» разрывает на куски!\n\nЭтот фильм мог бы превратиться в аналог «Перевозчика», если бы не куча метафор, порой абсолютно блистательных, и темы любви и смерти, пронизывающие его насквозь.\nОтношения киллера Чева с его боссом Карлито, его встреча с девушкой – чудесная иллюстрация всем нам ненавистной мысли о том, что наши страдания в этом мире никому не интересны, и что мы будем делать со своей жизнью – решать нам. Интуиция Чева, его стремление остаться с любимой (неважно, что блондинкой – уж какая есть!) – это умение доверять себе и попытка устроить жизнь по своему сценарию, а не по чужому. И это – лишь два примера в бесконечной череде драйва, ржача, метафор и неприкрытой жестокости последних часов жизни фактического мертвеца по имени Чев Челиос.\n\nНо главный фон «Адреналина» - конечно, драйв, - безудержное бешенство человека, который слетел с катушек и понял, что ему все позволено. Абсолютная свобода, невыносимая легкость бытия – тайная мечта каждого, кто чувствует, как на него давят сотни миллиметров ртутного столба и груз социальных обязательств. Наблюдать за тем, как мечется Чев Челиос в погоне за острыми ощущениями – сплошное удовольствие; волны эндорфинов, приятно щекочущих мозг. Этот кайф «Адреналин» дарит большими ложками, позволяет его не пробовать, а просто жрать.\n\nВыходя с сеанса, чувствуешь себя переполненным, и спокойно идти совершенно невозможно – нужно бежать, сломя голову, как Чев Челиос, вколовший себе полный шприц эпинефрина.\n«Адреналин» - из тех картин, у которых не бывает сиквела, как «Бойцовский клуб», «Карты, деньги, два ствола» или «Большой куш». Хорошо, стиснем зубы и пересмотрим дважды… Но Чев Челиос не из тех, кто забывается легко и быстро. Если продолжению не бывать, пусть снимут приквел!"})
print(response)
print(response.content.decode())


response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID" : 2, "GenreID1" : 1,
                                                                           "GenreID2" : 3, "GenreID3" : 5,
                                                                         "Title" : "Сокровище нации (2004)",
                                                                         "Description" : "Современные охотники за сокровищами во главе с Беном Франклином Гейтсом узнают старинную легенду о сокровище, спрятанном еще отцами-основателями американского государства Джорджем Вашингтоном, Бенджамином Франклином и Томасом Джефферсоном.\n\nУзнать тайну клада можно с помощью Декларации Независимости США, в которой зашифрована разгадка. Но помимо разгадывания изощренной головоломки, нашим героям предстоит сразиться с любителями легкой наживы, жаждущим погреть руки на достоянии республики…"})

print(response)
print(response.content.decode())

print(response)
print(response.content.decode())
ContentID = int(response.content.decode())
Date = int(datetime.datetime.now().timestamp())
print(Date)

response = requests.post("http://127.0.0.1:8081/api/v1/reviews", json = {"Date" : Date, "UserID" : userIDs[randint(0, len(userIDs) - 1)], "ContentID" : ContentID, "Review" : "Довольно неплохой фильм, для всей семьи. Много интересного, много забавного, минимум крови, точнее сказать её вообще нет, (не считая момента с пальцем), полное отсутствия брутальных сцен, похабщины, и явно выраженного эмоционального напряжения на лице актёров, даже напротив, зло-герой смотрится довольно симпатично, и совсем не вызывает негатива.\nКороче говоря, это просто лёгкий приключенческий фильм, снятый именно так, как умеют это делать ребята из Диснея, которые пока не особо хотят уходить далеко от мультипликационного жанра. Приятного просмотра."})
print(response)
print(response.content.decode())

response = requests.post("http://127.0.0.1:8081/api/v1/reviews", json = {"Date" : Date, "UserID" : userIDs[randint(0, len(userIDs) - 1)], "ContentID" : ContentID, "Review" : "Секреты, загадки, тайны, массоны, погони, сокровища, схватки и, конечно же, любовь - все это, словно сумасшедший танец пьяного Тёртлтауба, смешивается в самый «трушный» приключенческий коктейль. Коктейль, который является почти что идеализированной «адвенчурой» для неискушенного данным жанром зрителя (каковым и является ваш покорный слуга, в данном случае) и настоящей историей, пропитанной духом прошлых времён.\nК слову, историчности в картине уделяют совсем не первое место. Но оно того и не нужно - за фактическими справками и точными датами, придирчивый зритель, для которого понятие искусства и документальности каким-то образом пересекаются, должен отправиться в ближайшую библиотеку и не смотреть данный художественный фильм.\nБеря в расчёт некоторые «проколы», можно сказать, что Николас Кейдж со своим образом несколько неординарного «помешанного», совсем не тот кладоискатель, который должен вскапывать гробницы и водить туда красавиц по типу Крюгер. Свою восьмизначную зарплату он забрал несколько халтуря, уступая в игре антагонисту Шона Бина и даже молодому подмастерью Гейтса, в исполнении смешного и миловидного Джастина Барта.\nС юмором тут, честно говоря, проблемы тоже есть. Но за той легкостью и безостановочной динамикой это играет только на руку всей работе, не перебивая главного и самого интересного стержня картины - загадки сокровищницы тамплиеров.\nОтталкиваясь от всего вышесказанного и «переваренного», можно констатировать, что «Сокровище нации» - прекрасный и завлекающий в своём жанре фильм, который был способен ещё на чуть более качественную реализацию, угадай Тёртлтауб с главным персонажем и убери несколько напыщенную интеллектом голову Гейтса.\n9 из 10"})
print(response)
print(response.content.decode())

# TV serias

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":4,"GenreID2":5,"GenreID3":4,"Title":"Кухня","Description":"«Во все времена еда была рядом с людьми. Она искушает, вдохновляет, восхищает, отвращает. Но человек не может без еды, а значит тот, кто создает еду, может управлять миром...». Так думал Максим Лавров, когда шел осуществлять свою мечту – устраиваться поваром в один из самых дорогих ресторанов столицы. Мечта сбылась, вот только на деле все вышло не так вкусно. Шеф-повар ресторана Виктор Баринов – настоящая звезда гастрономического бомонда, знающая как угодить взыскательной публике. Оборотная сторона этого идеального образа – злоупотребление алкоголем, страсть к азартным играм и невыносимый характер."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":1,"GenreID2":1,"GenreID3":1,"Title":"Во все тяжкие","Description":"Умирающий учитель химии начинает варить мет ради благополучия семьи. Выдающийся драматический сериал 2010-х"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":5,"GenreID2":5,"GenreID3":5,"Title":"Офис","Description":"Скучающие от безделья клерки пытаются ужиться с безумным боссом. Виртуозный ситком про рабочие будни"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":4,"GenreID2":1,"GenreID3":3,"Title":"17 мгновений весны","Description":"Несколько недель из жизни советского разведчика. Телесериал Татьяны Лиозновой с гениальным актерским составом"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":4,"GenreID2":4,"GenreID3":4,"Title":"Сопрано","Description":"Гангстер идет на терапию чтобы разобраться в себе. Сериал HBO, с которого началась золотая эра телевидения"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":1,"GenreID2":4,"GenreID3":7,"Title":"Игра Престолов","Description":"Рыцари, мертвецы и драконы — в эпической битве за судьбы мира. Сериал, который навсегда изменил телевидение"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":5,"GenreID2":5,"GenreID3":5,"Title":"Рик и Морти","Description":"Гениальный ученый втягивает внука в безумные авантюры. Выдающийся анимационный сериал Дэна Хармона"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":4,"GenreID2":5,"GenreID3":5,"Title":"Доктор Хаус","Description":"Диагност Грегори Хаус, работающий в больнице Принстон‑Плейнсборо, отличается несносным характером и пристрастием к опиатам. Зато он вместе с помощниками берётся лечить самые сложные заболевания, расследуя их причины словно настоящий сыщик."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":3,"GenreID2":4,"GenreID3":1,"Title":"Очень странные дела","Description":"В небольшом американском городе Хокинсе в середине 1980‑х пропадает мальчик Уилл Байерс. Мама верит, что ребёнок связывается с ней из параллельной реальности, и при поддержке местного шерифа пытается его отыскать. Тем временем друзья Уилла встречают юную Одиннадцатую. Сбежавшая из лаборатории девочка обладает сверхъестественными способностями."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":5,"GenreID2":5,"GenreID3":5,"Title":"В Филадельфии всегда солнечно","Description":"Четверо приятелей открывают ирландский бар в одном из неблагополучных районов Филадельфии. Причём герои относятся друг к другу не лучшим образом. Они не стесняются подставлять и обманывать товарищей, что часто приводит к самым забавным ситуациям."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":1,"GenreID2":4,"GenreID3":1,"Title":"Мандалорец","Description":"Опытный охотник за головами с планеты Мандалор должен выполнить очередное задание. Он отправляется на поиски неназванной цели, за которой гоняются могущественные силы. Но неожиданное открытие полностью изменяет жизнь героя, и перед ним возникают новые трудности.\n\n"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":1,"GenreID2":4,"GenreID3":4,"Title":"Острые козырьки","Description":"В английском городе Бирмингеме в начале 1920‑х одна за другой зарождаются мелкие банды. Среди них и «Острые козырьки» под предводительством братьев Шелби, участники которой вшивали лезвия в кепки."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":4,"GenreID2":7,"GenreID3":4,"Title":"Чёрное зеркало","Description":"Сериал‑антология рассказывает сатирические истории о влиянии технологий на нашу жизнь. Герои вырабатывают электроэнергию на тренажёрах, чтобы поучаствовать в телешоу, заказывают копии умерших близких, а то и просто спасаются от обезумевших машин."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":5,"GenreID2":5,"GenreID3":5,"Title":"Друзья","Description":"Сериал посвящён жизни шестерых близких друзей. Они проводят вместе почти всё свободное время, то общаясь в любимой кофейне, то занимаясь будничными делами дома. Герои помогают друг другу справляться с трудностями, ругаются, мирятся и, конечно, влюбляются."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":1,"GenreID2":3,"GenreID3":4,"Title":"Настоящий детектив","Description":"Три сезона мрачной антологии посвящены различным преступлениям. В первом напарники, спустя много лет после закрытия дела, возвращаются к поискам убийцы женщины. Во втором межведомственная группа расследует смерть чиновника. В последнем детективы ищут девочку после смерти её брата."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":3,"GenreID2":1,"GenreID3":4,"Title":"Шерлок","Description":"Ветеран войны Джон Ватсон выходит в отставку и знакомится с гениальным частным сыщиком Шерлоком Холмсом. У последнего очень дурной характер. Напарники распутывают сложные дела и начинают борьбу с самим королём преступного мира."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":4,"GenreID2":7,"GenreID3":7,"Title":"Чернобыль","Description":"В 1986 году на Чернобыльской АЭС происходит авария. На место катастрофы отправляются Валерий Легасов и Борис Щербина"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":5,"GenreID2":5,"GenreID3":5,"Title":"Футурама","Description":"Курьер-лузер переносится в XXXI век и учится жить в мире космолетов и роботов. Сатира от авторов «Симпсонов»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":5,"GenreID2":5,"GenreID3":5,"Title":"Кремниевая долина","Description":"Группа гиков живет в инкубаторе и двигает мир интернета вперед. Комедийный сериал о внутренней кухне стартапов"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":4,"GenreID2":4,"GenreID3":4,"Title":"Мастер и Маргарита","Description":"Нечистая сила в Москве середины 1930-х. Режиссер «Собачьего сердца» бережно экранизирует роман Булгакова"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":5,"GenreID2":5,"GenreID3":5,"Title":"Детство Шелдона","Description":"Юный вундеркинд поступает в среднюю школу. Спин-офф «Теории большого взрыва» о ранних годах Шелдона Купера"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":1,"GenreID2":3,"GenreID3":4,"Title":"Сверхъестественное","Description":"Братья Винчестеры стараются не выпустить демонов. Фантастический сериал, вдохновленный «Сумеречной зоной»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":3,"GenreID2":1,"GenreID3":4,"Title":"Обмани меня","Description":"Эксперт по лжи помогает полиции раскалывать преступников. Тим Рот в детективном сериале в духе «Доктора Хауса»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":4,"GenreID2":5,"GenreID3":1,"Title":"Доктор Кто","Description":"Пришелец, меняющий внешность, спасает Вселенную от космических злодеев. Научная фантастика с британским юмором"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":3,"GenreID1":4,"GenreID2":5,"GenreID3":3,"Title":"Форс-мажоры","Description":"Юрист-самоучка случайно попадает в амбициозную фирму. Стильное и острое производственное драмеди с Меган Маркл"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

# Books

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":9,"GenreID2":8,"GenreID3":10,"Title":"Маленький принц","Description":"«Небольшая история, в которой заключен огромный жизненный смысл. История, которая заставляет по-другому посмотреть на привычные вещи.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":10,"GenreID2":8,"GenreID3":4,"Title":"Собачье сердце","Description":"«Удивительно тонкая и саркастическая история о людях и их пороках. История об эксперименте, который доказал, что из животного можно сделать человека, а вот вывести «животное» из человека нельзя.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":11,"GenreID2":4,"GenreID3":4,"Title":"Над пропастью во ржи","Description":"«История подростка, который своими глазами показывает свое восприятие мира, точку мировоззрения, отречение от привычных принципов и устоев морали общества, которые не вписываются в его индивидуальные рамки.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":3,"GenreID2":4, "GenreID3":12,"Title":"Приключения Шерлока Холмса","Description":"«Легендарные расследования великого сыщика Шерлока, которые раскрывают подлость человеческой души. Истории, которые рассказывает друг и помощник»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":8,"GenreID2":4,"GenreID3":11,"Title":"Портрет Дориана Грея","Description":"«История о самолюбии, эгоизме и прочной душе. История, которая наглядно показывает, что может случиться с душой человека, мучаемой пороками.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":11,"GenreID2":4,"GenreID3":12,"Title":"Триумфальная арка","Description":"«После Первой Мировой войны множество эмигрантов оказались во Франции. В их числе и талантливый немецкий хирург Равик. Это история его жизни и любви на фоне пережитой войны.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":8,"GenreID2":4,"GenreID3":12,"Title":"Зеленая миля","Description":"«Пол Эджкомб бывший сотрудник тюрьмы, который служил в блоке для осужденных на смертную казнь. Он рассказывает историю жизни смертников, которым суждено было пройти Зеленую милю.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":11,"GenreID2":4,"GenreID3":13,"Title":"Робинзон Крузо","Description":"«Дневник моряка потерпевшего крушение корабля и прожившего в одиночестве на острове 28 лет. Ему пришлось пережить слишком много испытаний.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":11,"GenreID2":4,"GenreID3":7,"Title":"Мы","Description":"«Роман антиутопия, в которой описывается идеальное общество, где нет личного мнения, а все происходит по расписанию. Но даже в таком обществе найдется место вольнодумцам.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":11,"GenreID2":4,"GenreID3":12,"Title":"Повелитель мух","Description":"«Что будет с детьми, если они окажутся совсем одни? У детей тонкая натура, которая подвержена порокам довольно сильно. И милые добрые дети превращаются в настоящих чудовищ.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":11,"GenreID2":4,"GenreID3":9,"Title":"451 градус по Фаренгейту","Description":"«Книги наше будущее, а что будет, если их заменит ТВ и одно мнение? Ответ на этот вопрос дает писатель, который опередил свое время.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":11,"GenreID2":4,"GenreID3":9,"Title":"Преступление и наказание","Description":"«Книга из школьной программы, которую трудно понять в нежном подростковом возрасте. Писатель показал двойственность человеческой души, когда черное переплетается с белым. История о Раскольникове, который переживает внутреннюю борьбу.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":11,"GenreID2":4,"GenreID3":8,"Title":"Сто лет одиночества","Description":"«Здесь излишни слова. В этом романе жизнь каждого героя пронизана одиночеством, впрочем, как и городка, где живут эти люди.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":11,"GenreID2":4,"GenreID3":7,"Title":"Степной волк","Description":"«Книга о внутреннем кризисе, который может случиться с каждым. Внутреннее опустошение может погубить человека, если однажды на пути не встретиться человек, который даст в руки всего одну книгу…»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":11,"GenreID2":5,"GenreID3":13,"Title":"12 стульев","Description":"«Кто не знает Остапа Бендера и Кису Воробьянинова и их вечные неудачи, которые связаны с поиском злополучных брильянтов.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":11,"GenreID2":4,"GenreID3":6,"Title":"Великий Гэтсби","Description":"«Великолепный роман, который наполнен чувствами. На страницах книги ждет начало 20 века, когда люди были полны иллюзий и надежд. Эта история о жизненных ценностях и настоящей любви.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":11,"GenreID2":7,"GenreID3":12,"Title":"Бойцовский клуб","Description":"«Это история о людях, которые решили изменить этот грязный мир. История о человеке, который смог противостоять системе.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":11,"GenreID2":9,"GenreID3":4,"Title":"Отцы и дети","Description":"«Это роман о противоборстве двух поколений, о их различиях и непринятии идей. Идеи романа легко ложатся на современные реалии.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":11,"GenreID2":7,"GenreID3":4,"Title":"Отец Горио","Description":"«Удивительная история о безграничной и жертвенной любви отца к детям. О любви, которая никогда не была взаимной. О любви, которая погубила отца Горио.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":11,"GenreID2":12,"GenreID3":4,"Title":"Коллекционер","Description":"«Он простой служитель ратуши, одинокий и потерянный. У него есть страсть – коллекционирование бабочек. Но однажды он захотел к себе в коллекцию девушку, которая покорила его душу.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":11,"GenreID2":13,"GenreID3":4,"Title":"Моби Дик","Description":"«Ахав капитан китобойного судна поставил целью своей жизни — месть киту по имени Моби Дик. Вит погуби слишком много жизней, чтобы оставить его в живых. Но стоит капитану начать охоту, как на его корабле начинают происходить загадочные, а порой и страшные события.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":11,"GenreID2":6,"GenreID3":4,"Title":"Госпожа Бовари","Description":"«Эта история признана величайшим произведением мировой литературы. Эмма Бовари мечтает о красивой светской жизни, но ее супруг, провинциальный врач, не может удовлетворить её запросов. Она находит любовников, но смогут ли они исполнить мечту мадам Бовари?»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":3,"GenreID2":13,"GenreID3":4,"Title":"Облачный атлас","Description":"«История прошлого, настоящего и будущего. Истории людей из разных времен. Но эти истории составляют единую картину всего нашего мира.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":5,"GenreID2":13,"GenreID3":8,"Title":"Приключения Алисы в стране чудес","Description":"«Странная и загадочная история о девочке, которая в погоне за белым кроликом оказывается в другом и чудесатом мире.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":1,"GenreID1":5,"GenreID2":11,"GenreID3":8,"Title":"Понедельник начинается в субботу","Description":"«Фантастическая и увлекательная сказка, в которой магия переплетается с реальностью.»"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

# Videogames

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":4,"GenreID1":2,"GenreID2":1,"GenreID3":5,"Title":"Red Dead Redemtion 2","Description":"Датч, Артур, Билл, Хавьер и Мика сбегают из города на корабле, направляющемся на Кубу. Корабль попадает в шторм и тонет, а людей выбрасывает на берег острова Гуарма, где они оказываются втянутыми в войну между владельцами плантаций сахарного тростника и местным населением, обращённым в рабство."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":4,"GenreID1":3,"GenreID2":4,"GenreID3":1,"Title":"Detroit - become human","Description":"По сюжету игры, в недалёком будущем, в 2038 году, на Земле полным ходом идёт производство андроидов — роботов, созданных компанией «Киберлайф» для удовлетворения различных человеческих потребностей. Но могут ли эти роботы чувствовать?"})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":4,"GenreID1":1,"GenreID2":3,"GenreID3":5,"Title":"GTA V","Description":"Легендарная игра со свободным миром. Очередной шедевр от Rockstar."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":4,"GenreID1":1,"GenreID2":3,"GenreID3":5,"Title":"Ведьмак 3: Дикая Охота","Description":"Игра в стиле фэнтези, мир которой основан на славянской мифологии, повествует о ведьмаке Геральте из Ривии, охотнике на чудовищ, чья приёмная дочь Цири находится в опасности, будучи преследуемой Дикой Охотой — загадочной потусторонней силой, тайна которой раскрывается по ходу игры."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":4,"GenreID1":1,"GenreID2":3,"GenreID3":4,"Title":"Mafia II","Description":"События Mafia II разворачиваются в 1943—1951 годах в вымышленном американском городе Эмпайр-Бэй (на основе реального Нью-Йорка, Чикаго, Сан-Франциско, Лос-Анджелеса, Бостона и Детройта). В городе «заправляют» три основные мафиозные «семьи»: Винчи, Клементе и Фальконе."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":4,"GenreID1":1,"GenreID2":3,"GenreID3":4,"Title":"Half-Life 2","Description":"Half-Life представила миру молодого очкарика-физика, обладателя докторской степени, Гордона Фримена. Простой ассистент исследовательского центра Чёрная меса оказывается в центре космической войны, начавшейся из-за неудачного научного эксперимента с телепортацией."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":4,"GenreID1":1,"GenreID2":5,"GenreID3":4,"Title":"Far Cry 6","Description":"В 1967 году на Яре произошла революция, в результате которой был свергнут президент Габриэль Кастильо. Из-за этого Яра оказалась изолированной от мира на десятилетия, а её экономика — на грани коллапса.В 2014-м на выборах побеждает сын бывшего президента, Антон Кастильо. Он обещает создать в стране рай на земле с помощью торговли Вивиро — экспериментальным препаратом для лечения рака, который производят из местного табака. Для него, однако, нужно чрезвычайно токсичное удобрение, которое плохо влияет на здоровье рабочих.Семь лет спустя, Кастильо вводит «лотерею», по которой случайным образом выбираются рабочие для принудительного труда на плантациях. "})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":4,"GenreID1":1,"GenreID2":3,"GenreID3":4,"Title":"Assassin’s Creed Valhalla","Description":"Действие игры начинается в 873 году во времена завоеваний викингов. Игроку предстоит взять на себя роль викинга по имени Эйвор, ведущего своих сородичей от берегов холодной Норвегии до плодородных земель Англии в поисках нового дома. Клану Эйвор противостоят лидеры четырёх англо-саксонских королевств: Уэссекса, Нортумбрии, Восточной Англии и Мерсии во главе с Альфредом Великим. "})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":4,"GenreID1":1,"GenreID2":3,"GenreID3":4,"Title":"The Elder Scrolls V: Skyrim","Description":"События игры происходят спустя 200 лет после событий Oblivion в 201 году 4-й эры. Великая война между Империей и Альдмерским Доминионом закончилась 30 лет назад унизительным Конкордатом Белого Золота, ущемлявшим множество прав жителей Империи, в том числе и поклонение богу Талосу (Тайберу Септиму). Группировка националистически настроенных нордов, известных как «Братья Бури», во главе с ветераном войны Ульфриком Буревестником, ярлом Виндхельма, устраивает мятеж против власти Империи, вследствие чего в Скайриме начинается гражданская война. Дабы заявить о себе всему Скайриму, Ульфрик убивает верховного короля провинции Торуга с помощью практически легендарного драконьего Крика, но вскоре попадает в засаду имперцев. Между тем, никто даже не догадывается о гораздо более древней опасности, ожидающей всех впереди."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":4,"GenreID1":1,"GenreID2":4,"GenreID3":3,"Title":"The Last of Us","Description":"Действие основной части происходит спустя 20 лет после глобальной катастрофы — пандемии, вызванной мутировавшим грибом кордицепсом. Согласно сюжету, кордицепс попадает в человеческий организм в виде спор в результате вдыхания, через кровь или слюну заражённого. Затем гриб вызывает необратимые изменения в человеческом организме — жертва теряет разум и превращается в подобие зомби. Пандемия привела к упадку цивилизации и гибели большей части населения США и всей Земли. Часть выживших укрывается в изолированных карантинных зонах, охраняемых военными. Города за пределами карантинных зон превратились в опасные руины, населённые заражёнными, бандитами и каннибалами. Террористическая организация «Цикады» (в оригинале «Светлячки» от англ. Fireflies) пытается поднять население карантинных зон на восстание против военных, а также стремится разработать вакцину против инфекции."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":4,"GenreID1":3,"GenreID2":4,"GenreID3":3,"Title":"Heavy Rain","Description":"Отпраздновав день рождения своего сына Джейсона, Итан Марс со своей семьёй отправляются в торговый центр за покупками. Но поход по магазинам оборачивается печальными событиями — Джейсона и Итана сбивает машина. Мальчик умирает, а его отец впадает в шестимесячную кому. После комы Итан винит в смерти своего сына себя, разводится со своей женой Грейс, переезжает в небольшой загородный дом и испытывает психологическую травму, провалы в памяти и страдает агорафобией. Два года спустя, находясь со своим вторым сыном Шоном в парке аттракционов, Итан теряет сознание, а когда приходит в себя, то обнаруживает в своей руке только фигурку оригами."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":4,"GenreID1":1,"GenreID2":4,"GenreID3":3,"Title":"MASS EFFECT","Description":"Эта грандиозная космоопера в формате экшен РПГ от признанных мастеров из студии BioWare повествует о закате эпохи галактических цивилизаций, которые одна за другой становятся жертвами нашествия огромных кораблей под управлением ИИ, известных как Жнецы."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":4,"GenreID1":1,"GenreID2":3,"GenreID3":3,"Title":"Battlefield 3","Description":"Американский морпех, Генри Блэкбёрн, на допросе в ЦРУ рассказывает агентам о человеке по имени Соломон и угрозе, которую он представляет для США. Ему обещают помочь, но Блэкбёрн должен сказать, откуда он знает о Соломоне."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":4,"GenreID1":1,"GenreID2":4,"GenreID3":5,"Title":"CYBERPUNK 2077","Description":"История наемника Ви, который встрял в невероятную переделку и обзавелся чипом, на котором хранится личность одного из самых крутых рокеров и опасных террористов."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":4,"GenreID1":1,"GenreID2":3,"GenreID3":4,"Title":"Death Stranding","Description":"В основе геймплея лежит доставка посылок из одной точки мира в другую: управляя героем Сэмом, игрок должен выбрать заказ через терминал доставки одного из поселений, погрузить посылку на спину героя и добраться с ней до адресата — в среднем такой пеший поход занимает от пяти до пятнадцати минут. Сам акт путешествия задумывался испытанием и источником эмоций для игрока, наподобие туристического похода. Каждая успешная доставка оценивается по ряду параметров, и игрок получает от неигровых персонажей особые очки — «лайки», наподобие функции «Нравится» в социальных сетях; игрок получает больше «лайков», если посылка не будет повреждена при доставке, или, например, доставит больше предметов, чем нужно для минимально успешной доставки."})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

response = requests.post("http://127.0.0.1:8081/api/v1/content", json = {"TypeID":4,"GenreID1":1,"GenreID2":4,"GenreID3":3,"Title":"WOLFENSTEIN: THE NEW ORDER","Description":"Действие разворачивается в альтернативной вселенной, где нацисты победили во Второй мировой войне, а Сопротивление ведет освободительную борьбу с переменным успехом. "})
print(response)
print(response.content.decode())
contentIDs.append(int(response.content.decode()))

usercontent = {}

for _ in range(500):
  userId = randint(0, len(userIDs) - 1)
  contentId = randint(0, len(contentIDs) - 1)
  if userId in usercontent and contentId in usercontent[userId]:
    print("Conflict userID: {} contentID: {}".format(userId, contentId))
  else:
    if userId in usercontent:
      usercontent[userId].append(contentId)
    else:
      usercontent[userId] = []
      usercontent[userId].append(contentId)
    response = requests.post("http://127.0.0.1:8081/api/v1/likes", json = {"UserID" : userIDs[userId], "ObjectID" : contentIDs[contentId]})
    print(response)
    print(response.content.decode())