<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Student;

class ApiController extends Controller
{
    public function create(Request $request)
    {
    	$students = new Student();

    	$students->fname = $request->fname;
    	$students->lname = $request->lname;
    	$students->email = $request->email;
    	$students->password = $request->password;

    	$students->save();

    	return response()->json($students);
    }

    public function show()
    {
    	$students = Student::all();
    	return response()->json($students);
    }

    public function showStudent($id)
    {
        $student = Student::find($id);
        return response()->json($student);
    }

    public function updateStudent(Request $request, $id)
    {
        $student = Student::find($id);

        $student->fname = $request->fname;
        $student->lname = $request->lname;
        $student->email = $request->email;
        $student->password = $request->password;

        $student->save();

        return response()->json($student);
    }

    public function deleteStudent($id)
    {
        $student = Student::find($id);
        $student->delete();
        return response()->json($student);
    }

}
