namespace Leetcode.FloodFill;

public class Solution {
    public int[][] FloodFill(int[][] image, int sr, int sc, int newColor) {
        if (image[sr][sc] == newColor) return image;

        return FillPaint(
            image: image,
            sr: sr,
            sc: sc,
            ocolor: image[sr][sc],
            ncolor: newColor
        );
    }

    private int[][] FillPaint(int[][] image, int sr, int sc, int ocolor, int ncolor) {
        if (sr < 0 || sr > image.Length) return image;
        if (sc < 0 || sc > image[0].Length) return image;

        var updatedImage = image;
        updatedImage[sr][sc] = ncolor;

        var top = (col: sc, row: sr - 1);
        if (top.row >= 0 && image[top.row][top.col] == ocolor) {
            updatedImage[top.row][top.col] = ncolor;
            updatedImage = FillPaint(updatedImage, top.row, top.col, ocolor, ncolor);
        }

        var bot = (col: sc, row: sr + 1);
        if (bot.row < image.Length && image[bot.row][bot.col] == ocolor) {
            updatedImage[bot.row][bot.col] = ncolor;
            updatedImage = FillPaint(updatedImage, bot.row, bot.col, ocolor, ncolor);
        }

        var lef = (col: sc - 1, row: sr);
        if (lef.col >= 0 && image[lef.row][lef.col] == ocolor) {
            updatedImage[lef.row][lef.col] = ncolor;
            updatedImage = FillPaint(updatedImage, lef.row, lef.col, ocolor, ncolor);
        }

        var rig = (col: sc + 1, row: sr);
        if (rig.col < image[0].Length && image[rig.row][rig.col] == ocolor) {
            updatedImage[rig.row][rig.col] = ncolor;
            updatedImage = FillPaint(updatedImage, rig.row, rig.col, ocolor, ncolor);
        }

        return updatedImage;
    }
}
