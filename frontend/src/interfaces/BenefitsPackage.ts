export interface AdminInterface {
    ID?: number;
    Email?: string;
    Password?: string;
    FirstName?: string;
    LastName?: string;
    BirthDay?: Date;
}

export interface AirlineInterface {
    ID?: number;
    AirlineName?: string;
    Movie_poster?: string;
    Movie_information?: string;
    Movie_video?: string;
    Movie_length?: string;
}

export interface BenefitsInterface {
    id?: number;
    BenefitsName?: string;
    FlyingFrom?: string;
    GoingTo?: string;
    PointRequired?: number;
    Quantity?: number;
    Code?: string;
    Class?: string;
    Trip?: string;
    Type?: string;
    AirlineID?: number;
    AdminID?: number;
}

export interface MemberInterface {
    id?: number;
    Email?: string;
    Password?: string;
    LastName?: string;
    BirthDay?: Date;
    Gender?: string;
    TotalPoint?: number;
}

export interface PointCalInterface {
    ID?: number;
    FlyingFrom?: string;
    GoingTo?: string;
    Point?: number;
    Class?: string;
    AirlineID?: number;
}
