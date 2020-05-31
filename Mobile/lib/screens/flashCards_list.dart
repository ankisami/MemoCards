import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:memocards/widgets/flashCard_question.dart';

import '../constants.dart';

class FlashCardsList extends StatelessWidget {
  String title = "Anglais";
  String category = "Langues";
  String autor = "Vous";
  @override
  Widget build(BuildContext context) {
    var size = MediaQuery.of(context).size;

    return Scaffold(
      body: Column(
        children: <Widget>[
          Stack(
            children: <Widget>[
              Container(
                height: size.height * .30,
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
                                style: TextStyle(
                                    color: Colors.black, fontSize: 13),
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
                    ],
                  ),
                ),
              ),
            ],
          ),
          Padding(
            padding: EdgeInsets.only(top: 5),
            // padding: EdgeInsets.only(top: size.height * .4 - 30),
            child: Column(
              children: <Widget>[
                Align(
                  alignment: Alignment.topRight,
                  child: IconButton(
                    icon: Icon(Icons.help),
                    onPressed: () {},
                  ),
                ),
                CardListQuestion(
                    answer: "A cat",
                    isDoubleFace: true,
                    numberQuestion: 1,
                    question: "Un chat"),
                CardListQuestion(
                    answer: "Car",
                    isDoubleFace: false,
                    numberQuestion: 2,
                    question: "Voiture"),
                CardListQuestion(
                    answer: "Sun",
                    isDoubleFace: false,
                    numberQuestion: 3,
                    question: "Soleil"),
                CardListQuestion(
                    answer: "Hello / Hi",
                    isDoubleFace: true,
                    numberQuestion: 4,
                    question: "Bonjour"),
              ],
            ),
          )
        ],
      ),
    );
  }
}
