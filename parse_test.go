package svgpath

import (
	"log"
	"math"
	"testing"
)

func TestParseA(t *testing.T) {
	line := `M1650,447.3c-9.6-20.6-24.4-39.3-45.2-57.1c-56.8-48.5-172.3-145.9-230.1-194.2
	c-56.4-47.2-90.2-71.4-124.7-89.3c-63.7-33-121.9,0.5-143.8,13.1c-8.3,4.8-19.2,11.2-30.8,18c-18.8,11.1-38.2,22.5-50.9,29.5
	C958.3,204,908.9,226.1,864.4,239c-7.9,2.3-20.9,5.1-33.3,7.9c-9.8,2.1-19,4.1-25.5,5.8c-43.4,11.3-71.7,38-92.3,57.5
	c-2.6,2.5-5.1,4.8-7.5,7c-24.2,22.3-71.8,69.5-94.3,96.5c-44.6,53.6-19.2,100.5-4,128.5c1.1,2.1,2.2,4.1,3.3,6.1
	c2.2,4.1,5,9.2,7.9,14.6c9,16.4,21.4,38.9,25.6,51.8c16.2,49.7-8,73.2-27.4,92.1l-1.2,1.1c-31.4,30.6-84.5,29-105.8,26.9
	c-20.7-2-46.8-5.5-72.1-8.8c-4.5-0.6-8.9-1.2-13.3-1.7c-13.6-1.8-24.8-3.4-33.7-4.6c-25.4-3.6-38.1-5.4-56.2-4.3
	c-36.8,2.2-64.9,17.3-83.5,44.8c-29.9,44.3-25.2,108.5-15.2,138.6c12.1,36.3,35.3,63.6,68.9,81.3c32.5,17.1,64.9,20,81.4,20.2
	c66.2,0.7,180.7,0.3,272.7-0.1c29.2-0.1,55.8-0.2,77.2-0.3c60.3-0.1,153.4,0.7,243.5,1.5c72.6,0.6,141.2,1.2,190.3,1.3l25.5,0.1
	c17.9,0,33.7,0.1,47.8,0.1c67.8,0,96.8-0.9,140.4-6.2c36.8-4.5,69.7-20.1,97.9-46.3c18.6-17.3,35.1-39.1,50.5-66.8
	c14.3-25.8,31.7-72.2,43.3-102.9c1.8-4.8,3.4-9.2,4.9-13.1c20.6-54.2,59-166.1,75.7-223.9C1665.8,509.5,1663.8,477,1650,447.3z`

	path, err := NewFromPath(line)
	if err != nil {
		t.Fatal(err)
	}

	pathLength := path.GetTotalLength()
	log.Printf("[A] path length: %f", pathLength)
}

func TestParseB(t *testing.T) {
	line := `M1636.6,538.2c-16.6,57.3-54.7,168.5-75.2,222.3c-1.5,3.9-3.1,8.3-4.9,13.1c-11.3,30.2-28.4,75.7-42,100.2
	c-34.8,62.5-78.4,96.2-133.4,102.9c-51.4,6.3-82.6,6.2-185.7,6l-25.5-0.1c-49-0.1-117.5-0.7-190.1-1.3c-84-0.7-170.6-1.5-230.9-1.5
	c-4.4,0-8.7,0-12.8,0c-21.4,0-48,0.2-77.2,0.3c-92,0.4-206.4,0.8-272.4,0.1c-17.1-0.2-104-5.3-131.5-87.8
	c-9.5-28.4-11.8-84.8,12.8-121.1c14.9-22,37.8-34.1,68.1-36c16-1,27.3,0.6,52.1,4.2c9,1.3,20.2,2.9,34,4.7c4.4,0.6,8.9,1.2,13.3,1.7
	c25.5,3.3,51.8,6.8,72.8,8.9c30,3,85.7,2.7,121.7-32.5l1.2-1.1c19.7-19.1,52.6-51.1,32.5-112.7c-4.8-14.7-17.2-37.2-27.1-55.3
	c-2.9-5.3-5.7-10.3-7.8-14.3c-1.1-2-2.2-4.1-3.4-6.3c-14.4-26.6-34.2-63,1.8-106.2c20.1-24.1,64.8-69.2,92.4-94.6
	c2.5-2.3,5-4.7,7.7-7.2c20.1-19,45.2-42.7,83.6-52.7c6.2-1.6,15.2-3.6,24.8-5.7c12.8-2.8,26-5.7,34.6-8.2
	c45.9-13.3,96.6-35.9,164.3-73.3c12.9-7.1,32.4-18.6,51.3-29.7c11.5-6.8,22.3-13.1,30.6-17.9c32.9-18.9,76.8-37.5,124.6-12.7
	c33.1,17.2,65.9,40.7,121.1,86.9c57.7,48.3,173.1,145.6,229.9,194C1636.1,443.3,1651.2,487.9,1636.6,538.2z`

	path, err := NewFromPath(line)
	if err != nil {
		t.Fatal(err)
	}

	pathLength := path.GetTotalLength()
	log.Printf("[B] path length: %f", pathLength)
}

func TestParse(t *testing.T) {
	limeRock := `M1650,447.3c-9.6-20.6-24.4-39.3-45.2-57.1c-56.8-48.5-172.3-145.9-230.1-194.2
	c-56.4-47.2-90.2-71.4-124.7-89.3c-63.7-33-121.9,0.5-143.8,13.1c-8.3,4.8-19.2,11.2-30.8,18c-18.8,11.1-38.2,22.5-50.9,29.5
	C958.3,204,908.9,226.1,864.4,239c-7.9,2.3-20.9,5.1-33.3,7.9c-9.8,2.1-19,4.1-25.5,5.8c-43.4,11.3-71.7,38-92.3,57.5
	c-2.6,2.5-5.1,4.8-7.5,7c-24.2,22.3-71.8,69.5-94.3,96.5c-44.6,53.6-19.2,100.5-4,128.5c1.1,2.1,2.2,4.1,3.3,6.1
	c2.2,4.1,5,9.2,7.9,14.6c9,16.4,21.4,38.9,25.6,51.8c16.2,49.7-8,73.2-27.4,92.1l-1.2,1.1c-31.4,30.6-84.5,29-105.8,26.9
	c-20.7-2-46.8-5.5-72.1-8.8c-4.5-0.6-8.9-1.2-13.3-1.7c-13.6-1.8-24.8-3.4-33.7-4.6c-25.4-3.6-38.1-5.4-56.2-4.3
	c-36.8,2.2-64.9,17.3-83.5,44.8c-29.9,44.3-25.2,108.5-15.2,138.6c12.1,36.3,35.3,63.6,68.9,81.3c32.5,17.1,64.9,20,81.4,20.2
	c66.2,0.7,180.7,0.3,272.7-0.1c29.2-0.1,55.8-0.2,77.2-0.3c60.3-0.1,153.4,0.7,243.5,1.5c72.6,0.6,141.2,1.2,190.3,1.3l25.5,0.1
	c17.9,0,33.7,0.1,47.8,0.1c67.8,0,96.8-0.9,140.4-6.2c36.8-4.5,69.7-20.1,97.9-46.3c18.6-17.3,35.1-39.1,50.5-66.8
	c14.3-25.8,31.7-72.2,43.3-102.9c1.8-4.8,3.4-9.2,4.9-13.1c20.6-54.2,59-166.1,75.7-223.9C1665.8,509.5,1663.8,477,1650,447.3z
	 M1636.6,538.2c-16.6,57.3-54.7,168.5-75.2,222.3c-1.5,3.9-3.1,8.3-4.9,13.1c-11.3,30.2-28.4,75.7-42,100.2
	c-34.8,62.5-78.4,96.2-133.4,102.9c-51.4,6.3-82.6,6.2-185.7,6l-25.5-0.1c-49-0.1-117.5-0.7-190.1-1.3c-84-0.7-170.6-1.5-230.9-1.5
	c-4.4,0-8.7,0-12.8,0c-21.4,0-48,0.2-77.2,0.3c-92,0.4-206.4,0.8-272.4,0.1c-17.1-0.2-104-5.3-131.5-87.8
	c-9.5-28.4-11.8-84.8,12.8-121.1c14.9-22,37.8-34.1,68.1-36c16-1,27.3,0.6,52.1,4.2c9,1.3,20.2,2.9,34,4.7c4.4,0.6,8.9,1.2,13.3,1.7
	c25.5,3.3,51.8,6.8,72.8,8.9c30,3,85.7,2.7,121.7-32.5l1.2-1.1c19.7-19.1,52.6-51.1,32.5-112.7c-4.8-14.7-17.2-37.2-27.1-55.3
	c-2.9-5.3-5.7-10.3-7.8-14.3c-1.1-2-2.2-4.1-3.4-6.3c-14.4-26.6-34.2-63,1.8-106.2c20.1-24.1,64.8-69.2,92.4-94.6
	c2.5-2.3,5-4.7,7.7-7.2c20.1-19,45.2-42.7,83.6-52.7c6.2-1.6,15.2-3.6,24.8-5.7c12.8-2.8,26-5.7,34.6-8.2
	c45.9-13.3,96.6-35.9,164.3-73.3c12.9-7.1,32.4-18.6,51.3-29.7c11.5-6.8,22.3-13.1,30.6-17.9c32.9-18.9,76.8-37.5,124.6-12.7
	c33.1,17.2,65.9,40.7,121.1,86.9c57.7,48.3,173.1,145.6,229.9,194C1636.1,443.3,1651.2,487.9,1636.6,538.2z`

	path, err := NewFromPath(limeRock)
	if err != nil {
		t.Fatal(err)
	}

	pathLength := path.GetTotalLength()
	log.Printf("path length: %f", pathLength)
}

func TestParseValues(t *testing.T) {
	str := "-9.6-20.6-24.4-39.3-45.2-57.1"
	values, err := parseValues(str)
	if err != nil {
		t.Fatal(err)
	}

	if len(values) != 6 {
		t.Fatalf("expected 6 values, got %d", len(values))
	}

	if !inDelta(values[0], -9.6, epsilon) {
		t.Fatalf("expected -9.6, got %f", values[0])
	} else if !inDelta(values[1], -20.6, epsilon) {
		t.Fatalf("expected -20.6, got %f", values[1])
	} else if !inDelta(values[2], -24.4, epsilon) {
		t.Fatalf("expected -24.4, got %f", values[2])
	} else if !inDelta(values[3], -39.3, epsilon) {
		t.Fatalf("expected -39.3, got %f", values[3])
	} else if !inDelta(values[4], -45.2, epsilon) {
		t.Fatalf("expected -45.2, got %f", values[4])
	} else if !inDelta(values[5], -57.1, epsilon) {
		t.Fatalf("expected -57.1, got %f", values[5])
	}
}

func TestGetPartsSimplerPath(t *testing.T) {
	path, err := NewFromPath("m10,0l10,0")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	parts, err := path.GetParts()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(parts) != 1 {
		t.Fatalf("expected 1 part, got %d", len(parts))
	}
}

func TestGetPartsSimplePath(t *testing.T) {
	path, err := NewFromPath("m10,0l10,0l10,0")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	parts, err := path.GetParts()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(parts) != 2 {
		t.Fatalf("expected 2 parts, got %d", len(parts))
	}

	p0 := parts[0]
	if p0.Start.X != 10 {
		t.Fatalf("expected p0.Start.X to be 10, got %f (%#v)", p0.Start.X, p0)
	} else if p0.Start.Y != 0 {
		t.Fatalf("expected p0.Start.Y to be 0, got %f (%#v)", p0.Start.Y, p0)
	}

	p1 := parts[1]
	if p1.Start.X != 20 {
		t.Fatalf("expected p1.Start.X to be 20, got %f (%#v)", p1.Start.X, p0)
	} else if p1.Start.Y != 0 {
		t.Fatalf("expected p1.Start.Y to be 0, got %f (%#v)", p1.Start.Y, p0)
	}

	if p0.End.X != 20 {
		t.Fatalf("expected p0.End.X to be 20, got %f (%#v)", p0.End.X, p0)
	} else if p0.End.Y != 0 {
		t.Fatalf("expected p0.End.Y to be 0, got %f (%#v)", p0.End.Y, p0)
	}

	if p1.End.X != 30 {
		t.Fatalf("expected p1.End.X to be 30, got %f (%#v)", p1.End.X, p0)
	} else if p1.End.Y != 0 {
		t.Fatalf("expected p1.End.Y to be 0, got %f (%#v)", p1.End.Y, p0)
	}

	if p0.GetTotalLength() != 10 {
		t.Fatalf("expected p0.GetTotalLength() to be 10, got %f (%#v)", p0.GetTotalLength(), p0)
	} else if p1.GetTotalLength() != 10 {
		t.Fatalf("expected p1.GetTotalLength() to be 10, got %f (%#v)", p1.GetTotalLength(), p0)
	}

	p0pAt5, err := p0.GetPointAtLength(5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if p0pAt5.X != 15 {
		t.Fatalf("expected p0pAt5.X to be 15, got %f (%#v)", p0pAt5.X, p0pAt5)
	} else if p0pAt5.Y != 0 {
		t.Fatalf("expected p0pAt5.Y to be 0, got %f (%#v)", p0pAt5.Y, p0pAt5)
	}

	p1pAt5, err := p1.GetPointAtLength(5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if p1pAt5.X != 25 {
		t.Fatalf("expected p1pAt5.X to be 25, got %f (%#v)", p1pAt5.X, p1pAt5)
	} else if p1pAt5.Y != 0 {
		t.Fatalf("expected p1pAt5.Y to be 0, got %f (%#v)", p1pAt5.Y, p1pAt5)
	}

	p0tAt5, err := p0.GetTangentAtLength(5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if p0tAt5.X != 1 {
		t.Fatalf("expected p0tAt5.X to be 1, got %f (%#v)", p0tAt5.X, p0tAt5)
	} else if p0tAt5.Y != 0 {
		t.Fatalf("expected p0tAt5.Y to be 0, got %f (%#v)", p0tAt5.Y, p0tAt5)
	}

	p1tAt5, err := p1.GetTangentAtLength(5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if p1tAt5.X != 1 {
		t.Fatalf("expected p1tAt5.X to be 1, got %f (%#v)", p1tAt5.X, p1tAt5)
	} else if p1tAt5.Y != 0 {
		t.Fatalf("expected p1tAt5.Y to be 0, got %f (%#v)", p1tAt5.Y, p1tAt5)
	}

	p0propAt5, err := p0.GetPropertiesAtLength(5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if p0propAt5.X != 15 {
		t.Fatalf("expected p0propAt5.X to be 15, got %f (%#v)", p0propAt5.X, p0propAt5)
	} else if p0propAt5.Y != 0 {
		t.Fatalf("expected p0propAt5.Y to be 0, got %f (%#v)", p0propAt5.Y, p0propAt5)
	} else if p0propAt5.TangentX != 1 {
		t.Fatalf("expected p0propAt5.TangentX to be 1, got %f (%#v)", p0propAt5.TangentX, p0propAt5)
	} else if p0propAt5.TangentY != 0 {
		t.Fatalf("expected p0propAt5.TangentY to be 0, got %f (%#v)", p0propAt5.TangentY, p0propAt5)
	}

	p1propAt5, err := p1.GetPropertiesAtLength(5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if p1propAt5.X != 25 {
		t.Fatalf("expected p1propAt5.X to be 25, got %f (%#v)", p1propAt5.X, p1propAt5)
	} else if p1propAt5.Y != 0 {
		t.Fatalf("expected p1propAt5.Y to be 0, got %f (%#v)", p1propAt5.Y, p1propAt5)
	} else if p1propAt5.TangentX != 1 {
		t.Fatalf("expected p1propAt5.TangentX to be 1, got %f (%#v)", p1propAt5.TangentX, p1propAt5)
	} else if p1propAt5.TangentY != 0 {
		t.Fatalf("expected p1propAt5.TangentY to be 0, got %f (%#v)", p1propAt5.TangentY, p1propAt5)
	}
}

func TestGetPartsSimplePathDistances(t *testing.T) {
	path, err := NewFromPath("M100,200 C100,100 250,100 250,200 S400,300 400,200")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	parts, err := path.GetParts()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(parts) != 2 {
		t.Fatalf("expected 2 parts, got %d", len(parts))
	}

	p0 := parts[0]
	p0pAt5, err := p0.GetPointAtLength(5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	pAtp05, err := path.GetPointAtLength(5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if p0pAt5.X != pAtp05.X {
		t.Fatalf("expected p0pAt5.X to be %f, got %f", pAtp05.X, p0pAt5.X)
	} else if p0pAt5.Y != pAtp05.Y {
		t.Fatalf("expected p0pAt5.Y to be %f, got %f", pAtp05.Y, p0pAt5.Y)
	}

	p1 := parts[1]
	p1pAt5, err := p1.GetPointAtLength(5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	pAtp15, err := path.GetPointAtLength(p0.GetTotalLength() + 5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if p1pAt5.X != pAtp15.X {
		t.Fatalf("expected p1pAt5.X to be %f, got %f", pAtp15.X, p1pAt5.X)
	} else if p1pAt5.Y != pAtp15.Y {
		t.Fatalf("expected p1pAt5.Y to be %f, got %f", pAtp15.Y, p1pAt5.Y)
	}
}

func TestOverloadedMoveTo1(t *testing.T) {
	segments, err := Parse("m 12.5,52 39,0 0,-40 -39,0 z")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(segments) != 5 {
		t.Fatalf("expected 5 segments, got %d: %v", len(segments), segments)
	}

	s0 := segments[0]
	if s0.Command != 'm' {
		t.Fatalf("expected s0.Command to be 'm', got %c", s0.Command)
	} else if len(s0.Args) != 2 {
		log.Printf("s0: %#v", s0)
		t.Fatalf("expected s0.Args to have 2 values, got %d", len(s0.Args))
	} else if s0.Args[0] != 12.5 {
		t.Fatalf("expected s0.Args[0] to be 12.5, got %f", s0.Args[0])
	} else if s0.Args[1] != 52 {
		t.Fatalf("expected s0.Args[1] to be 52, got %f", s0.Args[1])
	}

	s1 := segments[1]
	if s1.Command != 'l' {
		t.Fatalf("expected s1.Command to be 'l', got %c", s1.Command)
	} else if len(s1.Args) != 2 {
		t.Fatalf("expected s1.Args to have 2 values, got %d", len(s1.Args))
	} else if s1.Args[0] != 39 {
		t.Fatalf("expected s1.Args[0] to be 39, got %f", s1.Args[0])
	} else if s1.Args[1] != 0 {
		t.Fatalf("expected s1.Args[1] to be 0, got %f", s1.Args[1])
	}

	s2 := segments[2]
	if s2.Command != 'l' {
		t.Fatalf("expected s2.Command to be 'l', got %c", s2.Command)
	} else if len(s2.Args) != 2 {
		t.Fatalf("expected s2.Args to have 2 values, got %d", len(s2.Args))
	} else if s2.Args[0] != 0 {
		t.Fatalf("expected s2.Args[0] to be 0, got %f", s2.Args[0])
	} else if s2.Args[1] != -40 {
		t.Fatalf("expected s2.Args[1] to be -40, got %f", s2.Args[1])
	}

	s3 := segments[3]
	if s3.Command != 'l' {
		t.Fatalf("expected s3.Command to be 'l', got %c", s3.Command)
	} else if len(s3.Args) != 2 {
		t.Fatalf("expected s3.Args to have 2 values, got %d", len(s3.Args))
	} else if s3.Args[0] != -39 {
		t.Fatalf("expected s3.Args[0] to be -39, got %f", s3.Args[0])
	} else if s3.Args[1] != 0 {
		t.Fatalf("expected s3.Args[1] to be 0, got %f", s3.Args[1])
	}

	s4 := segments[4]
	if s4.Command != 'z' {
		t.Fatalf("expected s4.Command to be 'z', got %c", s4.Command)
	} else if len(s4.Args) != 0 {
		t.Fatalf("expected s4.Args to have 0 values, got %d", len(s4.Args))
	}
}

func TestOverloadedMoveTo2(t *testing.T) {
	segments, err := Parse("M 12.5,52 39,0 0,-40 -39,0 z")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(segments) != 5 {
		t.Fatalf("expected 5 segments, got %d: %v", len(segments), segments)
	}

	s0 := segments[0]
	if s0.Command != 'M' {
		t.Fatalf("expected s0.Command to be 'M', got %c", s0.Command)
	} else if len(s0.Args) != 2 {
		log.Printf("s0: %#v", s0)
		t.Fatalf("expected s0.Args to have 2 values, got %d", len(s0.Args))
	} else if s0.Args[0] != 12.5 {
		t.Fatalf("expected s0.Args[0] to be 12.5, got %f", s0.Args[0])
	} else if s0.Args[1] != 52 {
		t.Fatalf("expected s0.Args[1] to be 52, got %f", s0.Args[1])
	}

	s1 := segments[1]
	if s1.Command != 'L' {
		t.Fatalf("expected s1.Command to be 'L', got %c", s1.Command)
	} else if len(s1.Args) != 2 {
		t.Fatalf("expected s1.Args to have 2 values, got %d", len(s1.Args))
	} else if s1.Args[0] != 39 {
		t.Fatalf("expected s1.Args[0] to be 39, got %f", s1.Args[0])
	} else if s1.Args[1] != 0 {
		t.Fatalf("expected s1.Args[1] to be 0, got %f", s1.Args[1])
	}

	s2 := segments[2]
	if s2.Command != 'L' {
		t.Fatalf("expected s2.Command to be 'L', got %c", s2.Command)
	} else if len(s2.Args) != 2 {
		t.Fatalf("expected s2.Args to have 2 values, got %d", len(s2.Args))
	} else if s2.Args[0] != 0 {
		t.Fatalf("expected s2.Args[0] to be 0, got %f", s2.Args[0])
	} else if s2.Args[1] != -40 {
		t.Fatalf("expected s2.Args[1] to be -40, got %f", s2.Args[1])
	}

	s3 := segments[3]
	if s3.Command != 'L' {
		t.Fatalf("expected s3.Command to be 'L', got %c", s3.Command)
	} else if len(s3.Args) != 2 {
		t.Fatalf("expected s3.Args to have 2 values, got %d", len(s3.Args))
	} else if s3.Args[0] != -39 {
		t.Fatalf("expected s3.Args[0] to be -39, got %f", s3.Args[0])
	} else if s3.Args[1] != 0 {
		t.Fatalf("expected s3.Args[1] to be 0, got %f", s3.Args[1])
	}

	s4 := segments[4]
	if s4.Command != 'z' {
		t.Fatalf("expected s4.Command to be 'z', got %c", s4.Command)
	} else if len(s4.Args) != 0 {
		t.Fatalf("expected s4.Args to have 0 values, got %d", len(s4.Args))
	}
}

func TestCurveTo(t *testing.T) {
	segmentsA, err := Parse("c 50,0 50,100 100,100 50,0 50,-100 100,-100")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(segmentsA) != 2 {
		t.Fatalf("expected 2 segments, got %d", len(segmentsA))
	}

	sA0 := segmentsA[0]
	if sA0.Command != 'c' {
		t.Fatalf("expected sA0.Command to be 'c', got %c", sA0.Command)
	} else if len(sA0.Args) != 6 {
		t.Fatalf("expected sA0.Args to have 6 values, got %d", len(sA0.Args))
	} else if !inDelta(sA0.Args[0], 50, epsilon) {
		t.Fatalf("expected sA0.Args[0] to be 50, got %f", sA0.Args[0])
	} else if !inDelta(sA0.Args[1], 0, epsilon) {
		t.Fatalf("expected sA0.Args[1] to be 0, got %f", sA0.Args[1])
	} else if !inDelta(sA0.Args[2], 50, epsilon) {
		t.Fatalf("expected sA0.Args[2] to be 50, got %f", sA0.Args[2])
	} else if !inDelta(sA0.Args[3], 100, epsilon) {
		t.Fatalf("expected sA0.Args[3] to be 100, got %f", sA0.Args[3])
	} else if !inDelta(sA0.Args[4], 100, epsilon) {
		t.Fatalf("expected sA0.Args[4] to be 100, got %f", sA0.Args[4])
	} else if !inDelta(sA0.Args[5], 100, epsilon) {
		t.Fatalf("expected sA0.Args[5] to be 100, got %f", sA0.Args[5])
	}

	sA1 := segmentsA[1]
	if sA1.Command != 'c' {
		t.Fatalf("expected sA1.Command to be 'c', got %c", sA1.Command)
	} else if len(sA1.Args) != 6 {
		t.Fatalf("expected sA1.Args to have 6 values, got %d", len(sA1.Args))
	} else if !inDelta(sA1.Args[0], 50, epsilon) {
		t.Fatalf("expected sA1.Args[0] to be 50, got %f", sA1.Args[0])
	} else if !inDelta(sA1.Args[1], 0, epsilon) {
		t.Fatalf("expected sA1.Args[1] to be 0, got %f", sA1.Args[1])
	} else if !inDelta(sA1.Args[2], 50, epsilon) {
		t.Fatalf("expected sA1.Args[2] to be 50, got %f", sA1.Args[2])
	} else if !inDelta(sA1.Args[3], -100, epsilon) {
		t.Fatalf("expected sA1.Args[3] to be -100, got %f", sA1.Args[3])
	} else if !inDelta(sA1.Args[4], 100, epsilon) {
		t.Fatalf("expected sA1.Args[4] to be 100, got %f", sA1.Args[4])
	} else if !inDelta(sA1.Args[5], -100, epsilon) {
		t.Fatalf("expected sA1.Args[5] to be -100, got %f", sA1.Args[5])
	}

	segmentsB, err := Parse("c 50,0 50,100 100,100 c 50,0 50,-100 100,-100")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(segmentsB) != 2 {
		t.Fatalf("expected 2 segments, got %d", len(segmentsA))
	}

	sA0 = segmentsB[0]
	if sA0.Command != 'c' {
		t.Fatalf("expected sA0.Command to be 'c', got %c", sA0.Command)
	} else if len(sA0.Args) != 6 {
		t.Fatalf("expected sA0.Args to have 6 values, got %d", len(sA0.Args))
	} else if !inDelta(sA0.Args[0], 50, epsilon) {
		t.Fatalf("expected sA0.Args[0] to be 50, got %f", sA0.Args[0])
	} else if !inDelta(sA0.Args[1], 0, epsilon) {
		t.Fatalf("expected sA0.Args[1] to be 0, got %f", sA0.Args[1])
	} else if !inDelta(sA0.Args[2], 50, epsilon) {
		t.Fatalf("expected sA0.Args[2] to be 50, got %f", sA0.Args[2])
	} else if !inDelta(sA0.Args[3], 100, epsilon) {
		t.Fatalf("expected sA0.Args[3] to be 100, got %f", sA0.Args[3])
	} else if !inDelta(sA0.Args[4], 100, epsilon) {
		t.Fatalf("expected sA0.Args[4] to be 100, got %f", sA0.Args[4])
	} else if !inDelta(sA0.Args[5], 100, epsilon) {
		t.Fatalf("expected sA0.Args[5] to be 100, got %f", sA0.Args[5])
	}

	sA1 = segmentsB[1]
	if sA1.Command != 'c' {
		t.Fatalf("expected sA1.Command to be 'c', got %c", sA1.Command)
	} else if len(sA1.Args) != 6 {
		t.Fatalf("expected sA1.Args to have 6 values, got %d", len(sA1.Args))
	} else if !inDelta(sA1.Args[0], 50, epsilon) {
		t.Fatalf("expected sA1.Args[0] to be 50, got %f", sA1.Args[0])
	} else if !inDelta(sA1.Args[1], 0, epsilon) {
		t.Fatalf("expected sA1.Args[1] to be 0, got %f", sA1.Args[1])
	} else if !inDelta(sA1.Args[2], 50, epsilon) {
		t.Fatalf("expected sA1.Args[2] to be 50, got %f", sA1.Args[2])
	} else if !inDelta(sA1.Args[3], -100, epsilon) {
		t.Fatalf("expected sA1.Args[3] to be -100, got %f", sA1.Args[3])
	} else if !inDelta(sA1.Args[4], 100, epsilon) {
		t.Fatalf("expected sA1.Args[4] to be 100, got %f", sA1.Args[4])
	} else if !inDelta(sA1.Args[5], -100, epsilon) {
		t.Fatalf("expected sA1.Args[5] to be -100, got %f", sA1.Args[5])
	}
}

func TestArcTo(t *testing.T) {
	segments, err := Parse("A 30 50 0 0 1 162.55 162.45")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(segments) != 1 {
		t.Fatalf("expected 1 segment, got %d", len(segments))
	}

	s0 := segments[0]
	if s0.Command != 'A' {
		t.Fatalf("expected s0.Command to be 'A', got %c", s0.Command)
	} else if len(s0.Args) != 7 {
		t.Fatalf("expected s0.Args to have 7 values, got %d", len(s0.Args))
	} else if !inDelta(s0.Args[0], 30, epsilon) {
		t.Fatalf("expected s0.Args[0] to be 30, got %f", s0.Args[0])
	} else if !inDelta(s0.Args[1], 50, epsilon) {
		t.Fatalf("expected s0.Args[1] to be 50, got %f", s0.Args[1])
	} else if !inDelta(s0.Args[2], 0, epsilon) {
		t.Fatalf("expected s0.Args[2] to be 0, got %f", s0.Args[2])
	} else if !inDelta(s0.Args[3], 0, epsilon) {
		t.Fatalf("expected s0.Args[3] to be 0, got %f", s0.Args[3])
	} else if !inDelta(s0.Args[4], 1, epsilon) {
		t.Fatalf("expected s0.Args[4] to be 1, got %f", s0.Args[4])
	} else if !inDelta(s0.Args[5], 162.55, epsilon) {
		t.Fatalf("expected s0.Args[5] to be 162.55, got %f", s0.Args[5])
	} else if !inDelta(s0.Args[6], 162.45, epsilon) {
		t.Fatalf("expected s0.Args[6] to be 162.45, got %f", s0.Args[6])
	}
}

func TestQuadraticCurveTo(t *testing.T) {
	segments, err := Parse("M10 80 Q 95 10 180 80")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(segments) != 2 {
		t.Fatalf("expected 2 segments, got %d", len(segments))
	}

	s0 := segments[0]
	if s0.Command != 'M' {
		t.Fatalf("expected s0.Command to be 'M', got %c", s0.Command)
	} else if len(s0.Args) != 2 {
		t.Fatalf("expected s0.Args to have 2 values, got %d", len(s0.Args))
	} else if !inDelta(s0.Args[0], 10, epsilon) {
		t.Fatalf("expected s0.Args[0] to be 10, got %f", s0.Args[0])
	} else if !inDelta(s0.Args[1], 80, epsilon) {
		t.Fatalf("expected s0.Args[1] to be 80, got %f", s0.Args[1])
	}

	s1 := segments[1]
	if s1.Command != 'Q' {
		t.Fatalf("expected s1.Command to be 'Q', got %c", s1.Command)
	} else if len(s1.Args) != 4 {
		t.Fatalf("expected s1.Args to have 4 values, got %d", len(s1.Args))
	} else if !inDelta(s1.Args[0], 95, epsilon) {
		t.Fatalf("expected s1.Args[0] to be 95, got %f", s1.Args[0])
	} else if !inDelta(s1.Args[1], 10, epsilon) {
		t.Fatalf("expected s1.Args[1] to be 10, got %f", s1.Args[1])
	} else if !inDelta(s1.Args[2], 180, epsilon) {
		t.Fatalf("expected s1.Args[2] to be 180, got %f", s1.Args[2])
	} else if !inDelta(s1.Args[3], 80, epsilon) {
		t.Fatalf("expected s1.Args[3] to be 80, got %f", s1.Args[3])
	}
}

func TestSmoothCurveTo(t *testing.T) {
	segments, err := Parse("S 1 2, 3 4")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(segments) != 1 {
		t.Fatalf("expected 1 segments, got %d", len(segments))
	}

	s0 := segments[0]
	if s0.Command != 'S' {
		t.Fatalf("expected s0.Command to be 'S', got %c", s0.Command)
	} else if len(s0.Args) != 4 {
		t.Fatalf("expected s0.Args to have 4 values, got %d", len(s0.Args))
	} else if !inDelta(s0.Args[0], 1, epsilon) {
		t.Fatalf("expected s0.Args[0] to be 1, got %f", s0.Args[0])
	} else if !inDelta(s0.Args[1], 2, epsilon) {
		t.Fatalf("expected s0.Args[1] to be 2, got %f", s0.Args[1])
	} else if !inDelta(s0.Args[2], 3, epsilon) {
		t.Fatalf("expected s0.Args[2] to be 3, got %f", s0.Args[2])
	} else if !inDelta(s0.Args[3], 4, epsilon) {
		t.Fatalf("expected s0.Args[3] to be 4, got %f", s0.Args[3])
	}
}

func TestSmoothQuadraticCurveTo(t *testing.T) {
	// test.deepEqual(parse("T 1 -2e2"), [["T", 1, -2e2]]);
	segments, err := Parse("T 1 -2e2")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(segments) != 1 {
		t.Fatalf("expected 1 segments, got %d", len(segments))
	}

	s0 := segments[0]
	if s0.Command != 'T' {
		t.Fatalf("expected s0.Command to be 'T', got %c", s0.Command)
	} else if len(s0.Args) != 2 {
		t.Fatalf("expected s0.Args to have 2 values, got %d", len(s0.Args))
	} else if !inDelta(s0.Args[0], 1, epsilon) {
		t.Fatalf("expected s0.Args[0] to be 1, got %f", s0.Args[0])
	} else if !inDelta(s0.Args[1], -2e2, epsilon) {
		t.Fatalf("expected s0.Args[1] to be -2e2, got %f", s0.Args[1])
	}
}

func TestInvalidT(t *testing.T) {
	_, err := Parse("t 1 2 3")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestInvalidEmptyString(t *testing.T) {
	_, err := Parse("")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

const epsilon = 0.000001

func inDelta(a float64, b float64, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}
