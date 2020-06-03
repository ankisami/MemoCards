import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';

import '../widgets/SearchBar.dart';
import '../widgets/bottom_nav_bar.dart';
import '../widgets/category_card.dart';
import 'deckDetails_screen.dart';

class HomeScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    var size = MediaQuery.of(context)
        .size; //this gonna give us total height and with of our device
    return Scaffold(
      bottomNavigationBar: BottomNavBar(),
      body: Stack(
        children: <Widget>[
          Container(
            // Here the height of the container is 45% of our total height
            height: size.height * .45,
            decoration: BoxDecoration(
                color: Color(0xFFF5CEB8),
                image: DecorationImage(
                  alignment: Alignment.centerLeft,
                  image: AssetImage("assets/images/undraw_pilates_gpdb.png"),
                )),
          ),
          SafeArea(
            child: Padding(
              padding: const EdgeInsets.symmetric(horizontal: 20),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: <Widget>[
                  Align(
                    alignment: Alignment.topRight,
                    child: Container(
                      alignment: Alignment.center,
                      height: 52,
                      width: 52,
                      decoration: BoxDecoration(
                        color: Color(0xFFF2BA1),
                        shape: BoxShape.circle,
                      ),
                      child: SvgPicture.asset("assets/icons/menu.svg"),
                    ),
                  ),
                  Text(
                    "Apprendre avec \nMemocards",
                    style: Theme.of(context)
                        .textTheme
                        .display1
                        .copyWith(fontWeight: FontWeight.w900),
                  ),
                  SearchBar(),
                  Expanded(
                    child: GridView.count(
                      crossAxisCount: 2,
                      childAspectRatio: 0.85,
                      crossAxisSpacing: 20,
                      mainAxisSpacing: 20,
                      children: <Widget>[
                        CategoryCard(
                            category: "Langues",
                            title: "Allemand",
                            svgSrc: "assets/icons/flags/germanFlag.svg",
                            press: () {}),
                        CategoryCard(
                            category: "Langues",
                            title: "Anglais",
                            svgSrc: "assets/icons/flags/englishFlag.svg",
                            press: () {
                              Navigator.push(
                                context,
                                MaterialPageRoute(builder: (context) {
                                  return DeckDetailsScreen();
                                }),
                              );
                            }),
                        CategoryCard(
                            category: "Quizz",
                            title: "Culture générale",
                            svgSrc: "assets/icons/quizz.svg",
                            press: () {}),
                        CategoryCard(
                            category: "Histoire/Géographie",
                            title: "Seconde Guerre Mondiale",
                            svgSrc: "assets/icons/education.svg",
                            press: () {}),
                        CategoryCard(
                            category: "Langues",
                            title: "Japonais",
                            svgSrc: "assets/icons/flags/japan.svg",
                            press: () {}),
                      ],
                    ),
                  )
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}
