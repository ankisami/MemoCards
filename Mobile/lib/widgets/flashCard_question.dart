import 'package:flutter/material.dart';
import 'package:memocards/constants.dart';

class CardListQuestion extends StatelessWidget {
  final int numberQuestion;
  final String question;
  final String answer;
  final bool isDoubleFace;
  const CardListQuestion({
    Key key,
    this.numberQuestion,
    this.question,
    this.answer,
    this.isDoubleFace,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    var size = MediaQuery.of(context).size;
    return Container(
      margin: EdgeInsets.symmetric(vertical: 8),
      padding: EdgeInsets.symmetric(vertical: 20, horizontal: 30),
      width: size.width - 48,
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(38.5),
        boxShadow: [
          BoxShadow(
            offset: Offset(0, 10),
            blurRadius: 33,
            color: Color(0xFFD3D3D3).withOpacity(.84),
          ),
        ],
      ),
      child: Row(
        children: <Widget>[
          RichText(
            text: TextSpan(
              children: [
                TextSpan(
                  text: "Question n°$numberQuestion\n",
                  style: TextStyle(color: kLightBlackColor),
                ),
                TextSpan(
                  text: "$question\n",
                  style: TextStyle(
                    fontSize: 16,
                    color: kBlackColor,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                TextSpan(
                  text: "$answer",
                  style: TextStyle(
                    fontSize: 16,
                    color: kBlackColor,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ],
            ),
          ),
          Spacer(),
          IconButton(
            icon: Icon(Icons.crop_rotate),
            color: isDoubleFace ? Colors.orange : Colors.blueGrey,
            onPressed: () {},
          ),
          IconButton(
            icon: Icon(Icons.mode_edit),
            onPressed: () {},
          ),
          IconButton(
            icon: Icon(Icons.delete),
            onPressed: () {},
          ),
        ],
      ),
    );
  }
}
