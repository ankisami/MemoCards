import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:memocards/constants.dart';
import 'package:memocards/screens/flashCards_details.dart';
import 'package:memocards/screens/flashCards_list.dart';
import 'package:memocards/widgets/option_deck_details.dart';
import 'package:memocards/widgets/start_game_bar.dart';

class DeckDetailsScreen extends StatelessWidget {
  String title = "Anglais";
  String description =
      "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book";
  String category = "Langues";
  String autor = "Vous";

  @override
  Widget build(BuildContext context) {
    var size = MediaQuery.of(context).size;

    return Scaffold(
      body: Stack(
        children: <Widget>[
          Container(
            height: size.height * .45,
            width: size.width,
            decoration: BoxDecoration(color: kBlueLightColor),
          ),
          SafeArea(
            child: Padding(
              padding: const EdgeInsets.symmetric(horizontal: 20.0),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: <Widget>[
                  SizedBox(
                    height: size.height * 0.05,
                  ),
                  Row(
                    children: <Widget>[
                      Column(
                        children: <Widget>[
                          Text(
                            title,
                            style: Theme.of(context)
                                .textTheme
                                .display1
                                .copyWith(fontWeight: FontWeight.w900),
                          ),
                          SizedBox(height: 6),
                          Text(
                            "Catégorie : $category\nAuteur : $autor",
                            style: TextStyle(color: Colors.black, fontSize: 13),
                          ),
                        ],
                      ),
                      Spacer(),
                      SvgPicture.asset(
                        "assets/icons/flags/englishFlag.svg",
                        width: 150,
                        height: 150,
                        alignment: Alignment.topRight,
                      ),
                    ],
                  ),
                  SizedBox(height: 0),
                  Text(
                    "Révision 10-20 MIN",
                    style: TextStyle(fontWeight: FontWeight.bold, fontSize: 20),
                  ),
                  SizedBox(height: 10),
                  SizedBox(
                    width: size.width * .9,
                    child: Text(
                      "$description",
                      textAlign: TextAlign.justify,
                    ),
                  ),
                  SizedBox(height: 23),
                  Wrap(
                    spacing: 20,
                    runSpacing: 30,
                    children: <Widget>[
                      StartGameBar(
                        press: () {
                          Navigator.push(
                            context,
                            MaterialPageRoute(builder: (context) {
                              return FlashCardsDetails();
                            }),
                          );
                        },
                      ),
                      OptionDecksDetails(
                        title: "Statistiques",
                        iconData: Icons.trending_up,
                        //iconData: Icons.poll,
                        color: Colors.blue[100],
                      ),
                      OptionDecksDetails(
                        title: "Modifier",
                        iconData: Icons.settings,
                        //iconData: Icons.settings,
                        color: Colors.orange[100],
                      ),
                      OptionDecksDetails(
                        title: "Historique",
                        iconData: Icons.calendar_today,
                        color: Colors.red[100],
                      ),
                      OptionDecksDetails(
                          title: "FlashCards",
                          iconData: Icons.style,
                          color: Colors.green[100],
                          press: () {
                            Navigator.push(
                              context,
                              MaterialPageRoute(builder: (context) {
                                return FlashCardsList();
                              }),
                            );
                          }),
                    ],
                  ),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}
